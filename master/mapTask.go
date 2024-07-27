package master

import (
	"sync"

	"github.com/google/uuid"
)

type TaskStatus int

const (
	IDLE TaskStatus = iota
	INPROGRESS
	COMPLETE
	FAILED
)

type MapTask struct {
	ID          string
	mu          sync.RWMutex
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

func (t *MapTask) GetId() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.ID
}

func (t *MapTask) GetTaskFile() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.TaskFile
}

func (t *MapTask) GetWorkerId() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.WorkerID
}

func (t *MapTask) GetNumReduce() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.NumReduce
}

func (t *MapTask) GetTaskStatus() TaskStatus {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.TaskStatus
}

func (t *MapTask) SetTaskStatus(status TaskStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.TaskStatus = status
}

func (t *MapTask) GetOutputFiles() []string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.OutputFiles
}

func (t *MapTask) SetOutputFiles(files []string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.OutputFiles = files
}
