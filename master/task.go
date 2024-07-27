package master

import "github.com/google/uuid"

type TaskStatus int

const (
	IDLE TaskStatus = iota
	INPROGRESS
	COMPLETE
	FAILED
)

type MapTask struct {
	ID          string
	TaskFile    string
	WorkerID    string
	NumReduce   int
	TaskStatus  TaskStatus
	OutputFiles []string
}

func NewMapTask(taskFile string, workerId string, numReduce int) *MapTask {
	return &MapTask{
		ID:         uuid.NewString(),
		TaskFile:   taskFile,
		WorkerID:   workerId,
		TaskStatus: IDLE,
		NumReduce:  numReduce,
	}
}

type ReduceTask struct {
	ID         string
	TaskFiles  []*ReduceDataNodeInfo
	WorkerID   string
	TaskStatus TaskStatus
	OutputFile string
}

func NewReduceTask(taskFiles []*ReduceDataNodeInfo, workerId string) *ReduceTask {
	return &ReduceTask{
		ID:         uuid.NewString(),
		TaskFiles:  taskFiles,
		WorkerID:   workerId,
		TaskStatus: IDLE,
	}
}

type ReduceDataNodeInfo struct {
	Filename string
	NodeIP   string
}
