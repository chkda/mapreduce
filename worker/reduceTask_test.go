package worker

import (
	"testing"
)

func TestNewReduceTask(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "mr-task1-0.txt", NodeIP: "localhost:5001"},
		{Filename: "mr-task2-0.txt", NodeIP: "localhost:5002"},
	}
	taskId := "reduce-task-123"
	
	task := NewReduceTask(taskFiles, taskId)
	
	if task == nil {
		t.Fatal("NewReduceTask returned nil")
	}
	
	if task.GetId() != taskId {
		t.Errorf("Expected task ID '%s', got '%s'", taskId, task.GetId())
	}
	
	if task.GetTaskStatus() != IDLE {
		t.Errorf("Expected initial status IDLE, got %v", task.GetTaskStatus())
	}
	
	gotFiles := task.GetTaskFiles()
	if len(gotFiles) != len(taskFiles) {
		t.Errorf("Expected %d task files, got %d", len(taskFiles), len(gotFiles))
	}
}

func TestReduceTask_StatusTransitions(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "test.txt", NodeIP: "localhost:5001"},
	}
	task := NewReduceTask(taskFiles, "task-1")
	
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
}

func TestReduceTask_OutputFile(t *testing.T) {
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "test.txt", NodeIP: "localhost:5001"},
	}
	task := NewReduceTask(taskFiles, "task-1")
	
	outputFile := "mr-out-task1.txt"
	task.SetOutputFile(outputFile)
	
	if task.GetOutputFile() != outputFile {
		t.Errorf("Expected output file '%s', got '%s'", outputFile, task.GetOutputFile())
	}
}
