package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

// GetArgs struct represents the arguments for the Get RPC.
type GetArgs struct {
	WorkerID string
}

// GetResp struct represents the response for the Get RPC.
type GetResp struct {
	Task         Task // Task information to be assigned to the worker.
	NumOfReduces int  // Number of reduce tasks in the job.
}

// FinishResp struct represents the response for the Finish RPC.
type FinishResp struct {
	Flag bool // Indicates whether the task completion was successful.
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}


