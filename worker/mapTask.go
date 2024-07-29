package worker

import "sync"

type TaskStatus int

const (
	IDLE TaskStatus = iota
	INPROGRESS
	COMPLETED
	FAILED
)

type MapTask struct {
	ID          string
	mu          sync.RWMutex
	TaskFile    string
	NumReduce   int
	TaskStatus  TaskStatus
	OutputFiles []string
}

func NewMapTask(taskFile string, taskId string, numReduce int) *MapTask {
	return &MapTask{
		ID:         taskId,
		TaskFile:   taskFile,
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
