package master

import (
	"testing"
)

func TestNewReduceTask(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "mr-task1-0.txt", NodeIP: "localhost:5001"},
		{Filename: "mr-task2-0.txt", NodeIP: "localhost:5002"},
	}
	workerId := "worker-123"
	
	task := NewReduceTask(taskFiles, workerId)
	
	if task == nil {
		t.Fatal("NewReduceTask returned nil")
	}
	
	if task.GetId() == "" {
		t.Error("Task ID should not be empty")
	}
	
	if task.GetWorkerId() != workerId {
		t.Errorf("Expected WorkerId '%s', got '%s'", workerId, task.GetWorkerId())
	}
	
	if task.GetTaskStatus() != IDLE {
		t.Errorf("Expected initial status IDLE, got %v", task.GetTaskStatus())
	}
	
	gotFiles := task.GetTaskFiles()
	if len(gotFiles) != len(taskFiles) {
		t.Errorf("Expected %d task files, got %d", len(taskFiles), len(gotFiles))
	}
}

func TestReduceTask_SetTaskStatus(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "test.txt", NodeIP: "localhost:5001"},
	}
	task := NewReduceTask(taskFiles, "worker-1")
	
	statuses := []TaskStatus{IDLE, INPROGRESS, COMPLETE, FAILED}
	
	for _, status := range statuses {
		task.SetTaskStatus(status)
		if task.GetTaskStatus() != status {
			t.Errorf("Expected status %v, got %v", status, task.GetTaskStatus())
		}
	}
}

func TestReduceTask_SetOutputFile(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "test.txt", NodeIP: "localhost:5001"},
	}
	task := NewReduceTask(taskFiles, "worker-1")
	
	outputFile := "mr-out-task123.txt"
	task.SetOutputFile(outputFile)
	
	if task.GetOutputFile() != outputFile {
		t.Errorf("Expected output file '%s', got '%s'", outputFile, task.GetOutputFile())
	}
}

func TestReduceTask_UniqueIDs(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "test.txt", NodeIP: "localhost:5001"},
	}
	
	task1 := NewReduceTask(taskFiles, "worker-1")
	task2 := NewReduceTask(taskFiles, "worker-2")
	task3 := NewReduceTask(taskFiles, "worker-3")
	
	ids := map[string]bool{
		task1.GetId(): true,
		task2.GetId(): true,
		task3.GetId(): true,
	}
	
	if len(ids) != 3 {
		t.Error("Task IDs are not unique")
	}
}

func TestReduceTask_TaskFiles(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "mr-task1-0.txt", NodeIP: "localhost:5001"},
		{Filename: "mr-task2-0.txt", NodeIP: "localhost:5002"},
		{Filename: "mr-task3-0.txt", NodeIP: "localhost:5003"},
	}
	
	task := NewReduceTask(taskFiles, "worker-1")
	gotFiles := task.GetTaskFiles()
	
	if len(gotFiles) != 3 {
		t.Errorf("Expected 3 task files, got %d", len(gotFiles))
	}
	
	for i, file := range taskFiles {
		if gotFiles[i].Filename != file.Filename {
			t.Errorf("Expected filename '%s', got '%s'", file.Filename, gotFiles[i].Filename)
		}
		if gotFiles[i].NodeIP != file.NodeIP {
			t.Errorf("Expected NodeIP '%s', got '%s'", file.NodeIP, gotFiles[i].NodeIP)
		}
	}
}

func TestReduceTask_ThreadSafety(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "test.txt", NodeIP: "localhost:5001"},
	}
	task := NewReduceTask(taskFiles, "worker-1")
	
	done := make(chan bool)
	
	// Concurrent status updates
	go func() {
		for i := 0; i < 100; i++ {
			task.SetTaskStatus(INPROGRESS)
		}
		done <- true
	}()
	
	go func() {
		for i := 0; i < 100; i++ {
			_ = task.GetTaskStatus()
		}
		done <- true
	}()
	
	// Wait for both goroutines
	<-done
	<-done
}

func TestReduceDataNodeInfo(t *testing.T) {
	info := &ReduceDataNodeInfo{
		Filename: "mr-task1-0.txt",
		NodeIP:   "192.168.1.100:5001",
	}
	
	if info.Filename != "mr-task1-0.txt" {
		t.Errorf("Expected filename 'mr-task1-0.txt', got '%s'", info.Filename)
	}
	
	if info.NodeIP != "192.168.1.100:5001" {
		t.Errorf("Expected NodeIP '192.168.1.100:5001', got '%s'", info.NodeIP)
	}
}
