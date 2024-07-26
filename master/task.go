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
	ID         string
	TaskFile   string
	WorkerID   string
	TaskStatus TaskStatus
	Result     []string
}

func NewMapTask(taskFile string, workerId string) *MapTask {
	return &MapTask{
		ID:         uuid.New().String(),
		TaskFile:   taskFile,
		WorkerID:   workerId,
		TaskStatus: IDLE,
	}
}
