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
	TaskStatus  TaskStatus
	OutputFiles []string
}

func NewMapTask(taskFile string, workerId string) *MapTask {
	return &MapTask{
		ID:         uuid.NewString(),
		TaskFile:   taskFile,
		WorkerID:   workerId,
		TaskStatus: IDLE,
	}
}

type ReduceTask struct {
	ID         string
	TaskFiles  []string
	WorkerID   string
	TaskStatus TaskStatus
	OutputFile string
}

func NewReduceTask(taskFiles []string, workerId string) *ReduceTask {
	return &ReduceTask{
		ID:         uuid.NewString(),
		TaskFiles:  taskFiles,
		WorkerID:   workerId,
		TaskStatus: IDLE,
	}
}
