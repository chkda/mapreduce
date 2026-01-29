package worker

import (
	"testing"
)

func TestNewMapTask(t *testing.T) {
	taskFile := "test-file.txt"
	taskId := "task-123"
	numReduce := 5
	
	task := NewMapTask(taskFile, taskId, numReduce)
	
	if task == nil {
		t.Fatal("NewMapTask returned nil")
	}
	
	if task.GetId() != taskId {
		t.Errorf("Expected task ID '%s', got '%s'", taskId, task.GetId())
	}
	
	if task.GetTaskFile() != taskFile {
		t.Errorf("Expected TaskFile '%s', got '%s'", taskFile, task.GetTaskFile())
	}
	
	if task.GetNumReduce() != numReduce {
		t.Errorf("Expected NumReduce %d, got %d", numReduce, task.GetNumReduce())
	}
	
	if task.GetTaskStatus() != IDLE {
		t.Errorf("Expected initial status IDLE, got %v", task.GetTaskStatus())
	}
}

func TestMapTask_StatusTransitions(t *testing.T) {
	task := NewMapTask("file.txt", "task-1", 3)
	
	// IDLE -> INPROGRESS
	task.SetTaskStatus(INPROGRESS)
	if task.GetTaskStatus() != INPROGRESS {
		t.Error("Failed to transition to INPROGRESS")
	}
	
	// INPROGRESS -> COMPLETED
	task.SetTaskStatus(COMPLETED)
	if task.GetTaskStatus() != COMPLETED {
		t.Error("Failed to transition to COMPLETED")
	}
	
	// Can also transition to FAILED
	task.SetTaskStatus(FAILED)
	if task.GetTaskStatus() != FAILED {
		t.Error("Failed to transition to FAILED")
	}
}

func TestMapTask_OutputFiles(t *testing.T) {
	task := NewMapTask("file.txt", "task-1", 3)
	
	outputFiles := []string{"mr-task1-0.txt", "mr-task1-1.txt", "mr-task1-2.txt"}
	task.SetOutputFiles(outputFiles)
	
	gotFiles := task.GetOutputFiles()
	if len(gotFiles) != len(outputFiles) {
		t.Errorf("Expected %d files, got %d", len(outputFiles), len(gotFiles))
	}
	
	for i, file := range outputFiles {
		if gotFiles[i] != file {
			t.Errorf("File[%d]: expected '%s', got '%s'", i, file, gotFiles[i])
		}
	}
}
