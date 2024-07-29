package worker

import "sync"

type ReduceTask struct {
	ID         string
	mu         sync.RWMutex
	TaskFiles  []*ReduceDataNodeInfo
	TaskStatus TaskStatus
	OutputFile string
}

func NewReduceTask(taskFiles []*ReduceDataNodeInfo, taskId string) *ReduceTask {
	return &ReduceTask{
		ID:         taskId,
		TaskFiles:  taskFiles,
		TaskStatus: IDLE,
	}
}

func (t *ReduceTask) GetId() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.ID
}

func (t *ReduceTask) GetTaskFiles() []*ReduceDataNodeInfo {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.TaskFiles
}

func (t *ReduceTask) GetTaskStatus() TaskStatus {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.TaskStatus
}

func (t *ReduceTask) SetTaskStatus(status TaskStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.TaskStatus = status
}

func (t *ReduceTask) SetOutputFile(file string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.OutputFile = file
}

func (t *ReduceTask) GetOutputFile() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.OutputFile
}

type ReduceDataNodeInfo struct {
	Filename string
	NodeIP   string
}
