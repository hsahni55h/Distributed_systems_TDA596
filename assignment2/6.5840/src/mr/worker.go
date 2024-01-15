package mr

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"sort"
	"strings"
	"time"
)

// KeyValue represents a key-value pair.
type KeyValue struct {
	Key   string
	Value string
}

// KeyValuePairArray is an array of KeyValue.
type KeyValuePairArray []KeyValue

// Len returns the length of the KeyValuePairArray.
func (a KeyValuePairArray) Len() int { return len(a) }

// Swap swaps the elements at indices i and j.
func (a KeyValuePairArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less returns true if the element at index i is less than the element at index j.
func (a KeyValuePairArray) Less(i, j int) bool { return a[i].Key < a[j].Key }

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go calls this function.
//
// Worker is called by main/mrworker.go.
// Worker is called by main/mrworker.go to perform map and reduce tasks.
func Worker(mapf func(string, string) []KeyValue, reducef func(string, []string) string) {
	// Generate a unique worker ID using the process ID.
	id := fmt.Sprintf("w-%v", os.Getpid())

	// Continuously request tasks from the coordinator.
	for {
		// Get the next task, number of reduces, and status from the coordinator.
		task, nReduce, status := GetTask(id)
		// fmt.Println(task.Type, task.Index)

		// Exit the loop if no more tasks are available.
		if !status {
			break
		}

		// Handle different task states.
		switch task.Type {
		case TaskDone:
			// If the task is marked as DONE, sleep for a short duration and continue.
			time.Sleep(1 * time.Second)
			break
		case TaskWait:
			// If the task is in WAIT state, sleep for a shorter duration and continue.
			time.Sleep(10 * time.Millisecond)
			continue
		case TaskMap:
			// fmt.Println("Starting TaskMap")
			// If the task is a MAP task, process the input files and create intermediate files.
			var intermediate []KeyValuePairArray
			for i := 0; i < nReduce; i++ {
				intermediate = append(intermediate, []KeyValue{})
			}

			// Process each input file and generate key-value pairs using the provided map function.
			for _, filename := range task.Files {
				file, err := os.Open(filename)
				if err != nil {
					log.Fatalf("In TaskMap: cannot open %v", filename)
				}
				content, err := ioutil.ReadAll(file)
				if err != nil {
					log.Fatalf("cannot read %v", filename)
				}
				file.Close()

				// Apply the map function to generate key-value pairs.
				kva := mapf(filename, string(content))

				// Distribute key-value pairs into different partitions based on hash.
				for _, kv := range kva {
					r := ihash(kv.Key) % nReduce
					intermediate[r] = append(intermediate[r], kv)
				}
			}

			// Save the intermediate key-value pairs to temporary files.
			var filenames []string
			for i, kva := range intermediate {
				filename := fmt.Sprintf("mr-%v-%v.tmp", task.Index, i)
				filenames = append(filenames, filename)
				file, _ := os.Create(filename)
				enc := json.NewEncoder(file)
				for _, kv := range kva {
					enc.Encode(&kv)
				}
				file.Close()
			}

			// Notify the coordinator that the MAP task is finished.
			if SendFinish(task) {
				// Rename temporary files to remove the ".tmp" extension.
				for _, filename := range filenames {
					os.Rename(filename, strings.TrimSuffix(filename, ".tmp"))
				}
			}
			// fmt.Println("Map done.")
			continue
		case TaskReduce:
			// If the task is a REDUCE task, process the intermediate files and generate output files.
			var intermediate KeyValuePairArray
			for _, filename := range task.Files {
				file, err := os.Open(filename)
				if err != nil {
					log.Fatalf("In TaskReduce: cannot open %v %v", filename, err)
				}
				dec := json.NewDecoder(file)

				// Read key-value pairs from intermediate files.
				for {
					var kv KeyValue
					if err := dec.Decode(&kv); err != nil {
						break
					}
					intermediate = append(intermediate, kv)
				}
			}

			// Sort intermediate key-value pairs by key.
			sort.Sort(intermediate)

			// Create the output file for the REDUCE task.
			oname := fmt.Sprintf("mr-out-%v.tmp", task.Index)
			ofile, _ := os.Create(oname)

			// Aggregate values for each key and apply the reduce function.
			for i := 0; i < len(intermediate); {
				j := i + 1
				for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
					j++
				}
				words := []string{}
				for k := i; k < j; k++ {
					words = append(words, intermediate[k].Value)
				}
				output := reducef(intermediate[i].Key, words)
				fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)
				i = j
			}

			// Close the output file and notify the coordinator that the REDUCE task is finished.
			ofile.Close()
			if SendFinish(task) {
				// Rename the output file to remove the ".tmp" extension.
				os.Rename(oname, strings.TrimSuffix(oname, ".tmp"))
			}
		}
	}
}

// GetTask fetches the next task for a worker with the given ID from the coordinator.
// It returns the received task, the number of reduces, and a boolean indicating the success of the request.
func GetTask(workerID string) (Task, int, bool) {
	// Create arguments for the RPC call.
	args := GetArgs{WorkerID: workerID}
	// Initialize an empty response.
	resp := GetResp{}
	// Make the RPC call to the coordinator to get the next task.
	success := call("Coordinator.Get", &args, &resp)

	// Check if the RPC call was successful.
	if success {
		// Return the received task, the number of reduces, and a true flag indicating success.
		return resp.Task, resp.NumOfReduces, true
	} else {
		// Log a fatal error and return default values with a false flag indicating failure.
		log.Fatalln("GetTask RPC call failed")
		return Task{}, 0, false
	}
}

// SendFinish sends a finish signal to the coordinator for the given task.
// It returns true if the coordinator acknowledges the finish; otherwise, it returns false.
func SendFinish(t Task) bool {
	// Create a response struct to store the coordinator's reply.
	resp := FinishResp{}

	// Make an RPC call to inform the coordinator about task completion.
	success := call("Coordinator.Finish", &t, &resp)

	// Check if the RPC call was successful.
	if success {
		// Return the finish flag received from the coordinator.
		return resp.Flag
	}

	// If the RPC call was not successful, return false.
	return false
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
