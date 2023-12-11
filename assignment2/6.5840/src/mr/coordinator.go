package mr

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	// "os"
	"sync"
	"time"
)

// Task represents a unit of work in the MapReduce job.
type Task struct {
    Type  byte       // Type of the task: MAP, REDUCE, DONE, WAIT
    Index     int    // Index of the task
    WorkerID  string // ID of the worker assigned to the task
    State     byte   // State of the task: IDLE, RUNNING, FINISH
    Files     []string // Files associated with the task
    Timestamp int64  // Timestamp of when the task was assigned or last updated
}

const (
    TaskIdle    byte = 0
    TaskRunning byte = 1
    TaskFinish  byte = 2
)

const (
    TaskMap    byte = 0
    TaskReduce byte = 1
    TaskDone   byte = 2
    TaskWait   byte = 3
)

// DoneTask represents a finished task in the MapReduce job.
var DoneTask = Task{TaskDone, 0, "", TaskFinish, nil, 0}

// WaitTask represents a task that is waiting in the MapReduce job.
var WaitTask = Task{TaskWait, 0, "", TaskFinish, nil, 0}

// Coordinator manages the MapReduce job and assigns tasks to workers.
type Coordinator struct {
    // Your definitions here.
    mapTasks    []Task   // List of map tasks
    reduceTasks []Task   // List of reduce tasks
    mutex       sync.Mutex // Mutex to control access to shared data structures
}


// Your code here -- RPC handlers for the worker to call.

// Get assigns a task to a worker and returns the task information.
func (c *Coordinator) Get(args *GetArgs, resp *GetResp) error {
    // Lock to ensure mutual exclusion
    c.mutex.Lock()
    defer c.mutex.Unlock()

    mapDone := true
    resp.NumOfReduces = len(c.reduceTasks)

    // Check for available map tasks
    for i, t := range c.mapTasks {
        if t.State == TaskIdle {
            mapDone = false
            c.mapTasks[i].WorkerID = args.WorkerID
            c.mapTasks[i].State = TaskRunning
            c.mapTasks[i].Timestamp = time.Now().Unix()
            resp.Task = c.mapTasks[i]
            return nil
        } else if t.State == TaskRunning {
            mapDone = false
            if time.Now().Unix()-t.Timestamp > 10 {
                c.mapTasks[i].WorkerID = args.WorkerID
                c.mapTasks[i].Timestamp = time.Now().Unix()
                resp.Task = c.mapTasks[i]
                return nil
            }
        }
    }

    // Check for available reduce tasks if map tasks are done
    reduceDone := true
    if mapDone {
        for i, t := range c.reduceTasks {
            if t.State == TaskIdle {
                reduceDone = false
                c.reduceTasks[i].State = TaskRunning
                c.reduceTasks[i].WorkerID = args.WorkerID
                c.reduceTasks[i].Timestamp = time.Now().Unix()
                resp.Task = c.reduceTasks[i]
                return nil
            } else if t.State == TaskRunning {
                reduceDone = false
                if time.Now().Unix()-t.Timestamp > 10 {
                    c.reduceTasks[i].WorkerID = args.WorkerID
                    c.reduceTasks[i].Timestamp = time.Now().Unix()
                    resp.Task = c.reduceTasks[i]
                    return nil
                }
            }
        }

        // If all tasks are done, return a special DONE_TASK
        if reduceDone {
            resp.Task = DoneTask
            return nil
        }
    }

    // If no tasks are available, return a special WAIT_TASK
    resp.Task = WaitTask
    return nil
}


// Finish marks a task as finished based on the provided task information.
func (c *Coordinator) Finish(task *Task, resp *FinishResp) error {
    // Lock to ensure mutual exclusion
    c.mutex.Lock()
    defer c.mutex.Unlock()

    index := task.Index

    // Check if it's a map task
    if task.Type == TaskMap {
        if c.isValidTaskIndex(index, c.mapTasks) && c.isMatchingTask(task, c.mapTasks[index]) {
            resp.Flag = c.updateTaskStatus(index, TaskFinish, c.mapTasks)
            return nil
        }
    }

    // Check if it's a reduce task
    if task.Type == TaskReduce {
        if c.isValidTaskIndex(index, c.reduceTasks) && c.isMatchingTask(task, c.reduceTasks[index]) {
            resp.Flag = c.updateTaskStatus(index, TaskFinish, c.reduceTasks)
            return nil
        }
    }

    // If the task is not valid or not matching, set Flag to false
    resp.Flag = false
    return nil
}

// isValidTaskIndex checks if the index is valid for the given task type.
func (c *Coordinator) isValidTaskIndex(index int, tasks []Task) bool {
    return index >= 0 && index < len(tasks)
}

// isMatchingTask checks if the provided task matches the task at the given index.
func (c *Coordinator) isMatchingTask(task *Task, target Task) bool {
    return task.Index == target.Index && task.WorkerID == target.WorkerID && task.State == TaskRunning &&
        time.Now().Unix()-target.Timestamp <= 10
}

// updateTaskStatus updates the status of the task at the given index.
func (c *Coordinator) updateTaskStatus(index int, state byte, tasks []Task) bool {
    if tasks[index].State == TaskRunning {
        tasks[index].State = state
        return true
    }
    return false
}



//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	// sockname := coordinatorSock()
	// os.Remove(sockname)
	// l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
// Done checks if all map and reduce tasks are finished.
func (c *Coordinator) Done() bool {
    // Lock to ensure mutual exclusion
    c.mutex.Lock()
    defer c.mutex.Unlock()

    // Check if all map tasks are finished
    if !c.areTasksFinished(c.mapTasks) {
        return false
    }

    // Check if all reduce tasks are finished
    if !c.areTasksFinished(c.reduceTasks) {
        return false
    }

    return true
}

// areTasksFinished checks if all tasks in the given list are finished.
func (c *Coordinator) areTasksFinished(tasks []Task) bool {
    for _, t := range tasks {
        if t.State != TaskFinish {
            return false
        }
    }
    return true
}


//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//

// MakeCoordinator creates a Coordinator with map and reduce tasks.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
    c := Coordinator{}

    // Initialize map tasks
    for i, file := range files {
        c.mapTasks = append(c.mapTasks, Task{
            Type:  TaskMap,
            Index: i,
            State: TaskIdle,
            Files: []string{file},
        })
    }

    // Initialize reduce tasks
    for i := 0; i < nReduce; i++ {
        var fs []string
        for j := 0; j < len(files); j++ {
            kvName := fmt.Sprintf("mr-%v-%v", j, i)
            fs = append(fs, kvName)
        }
        c.reduceTasks = append(c.reduceTasks, Task{
            Type:  TaskReduce,
            Index: i,
            State: TaskIdle,
            Files: fs,
        })
    }

    // Start the server
    c.server()
    return &c
}
