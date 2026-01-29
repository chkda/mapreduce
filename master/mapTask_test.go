package master

import (
	"testing"
)

func TestNewMapTask(t *testing.T) {
	taskFile := "test-file.txt"
	workerId := "worker-123"
	numReduce := 5
	
	task := NewMapTask(taskFile, workerId, numReduce)
	
	if task == nil {
		t.Fatal("NewMapTask returned nil")
	}
	
	if task.GetId() == "" {
		t.Error("Task ID should not be empty")
	}
	
	if task.GetTaskFile() != taskFile {
		t.Errorf("Expected TaskFile '%s', got '%s'", taskFile, task.GetTaskFile())
	}
	
	if task.GetWorkerId() != workerId {
		t.Errorf("Expected WorkerId '%s', got '%s'", workerId, task.GetWorkerId())
	}
	
	if task.GetNumReduce() != numReduce {
		t.Errorf("Expected NumReduce %d, got %d", numReduce, task.GetNumReduce())
	}
	
	if task.GetTaskStatus() != IDLE {
		t.Errorf("Expected initial status IDLE, got %v", task.GetTaskStatus())
	}
	
	if len(task.GetOutputFiles()) != 0 {
		t.Errorf("Expected empty OutputFiles, got %d files", len(task.GetOutputFiles()))
	}
}

func TestMapTask_SetTaskStatus(t *testing.T) {
	task := NewMapTask("file.txt", "worker-1", 3)
	
	statuses := []TaskStatus{IDLE, INPROGRESS, COMPLETE, FAILED}
	
	for _, status := range statuses {
		task.SetTaskStatus(status)
		if task.GetTaskStatus() != status {
			t.Errorf("Expected status %v, got %v", status, task.GetTaskStatus())
		}
	}
}

func TestMapTask_SetOutputFiles(t *testing.T) {
	task := NewMapTask("file.txt", "worker-1", 3)
	
	outputFiles := []string{
		"mr-task-0.txt",
		"mr-task-1.txt",
		"mr-task-2.txt",
	}
	
	task.SetOutputFiles(outputFiles)
	
	gotFiles := task.GetOutputFiles()
	if len(gotFiles) != len(outputFiles) {
		t.Errorf("Expected %d output files, got %d", len(outputFiles), len(gotFiles))
	}
	
	for i, file := range outputFiles {
		if gotFiles[i] != file {
			t.Errorf("Expected file[%d] '%s', got '%s'", i, file, gotFiles[i])
		}
	}
}

func TestMapTask_UniqueIDs(t *testing.T) {
	// Create multiple tasks and ensure IDs are unique
	task1 := NewMapTask("file1.txt", "worker-1", 3)
	task2 := NewMapTask("file2.txt", "worker-2", 3)
	task3 := NewMapTask("file3.txt", "worker-3", 3)
	
	ids := map[string]bool{
		task1.GetId(): true,
		task2.GetId(): true,
		task3.GetId(): true,
	}
	
	if len(ids) != 3 {
		t.Error("Task IDs are not unique")
	}
}

func TestMapTask_ThreadSafety(t *testing.T) {
	task := NewMapTask("file.txt", "worker-1", 3)
	
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


