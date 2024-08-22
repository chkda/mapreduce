package maptask

import (
	"testing"

	"github.com/chkda/mapreduce/internal/status"
)

func TestNew(t *testing.T) {
	task := New()

	if task.taskId == "" {
		t.Errorf("Expected non empty taskId")
	}

	if task.taskStatus != status.IDLE {
		t.Errorf("Expected idle taskStatus")
	}

	if task.outputFiles == nil {
		t.Errorf("Expected non nil output files")
	}
}

func TestWithOptions(t *testing.T) {
	task := New(
		WithNumReduce(3),
		WithTaskFile("file.txt"),
		WithTaskId("task-1"),
		WithWorkerId("worker-1"),
	)

	if task.numReduce != 3 {
		t.Errorf("Expected numReduce as 3, got %d", task.numReduce)
	}

	if task.taskFile != "file.txt" {
		t.Errorf("Expected taskFile as 'file.txt', got %s", task.taskFile)
	}

	if task.taskId != "task-1" {
		t.Errorf("Expected taskId as 'task-1', got %s", task.taskId)
	}

	if task.workerId != "worker-1" {
		t.Errorf("Expected workerId as 'worker-1', got %s", task.workerId)
	}
}
