package model

import (
	"testing"
	"time"
)

func TestNewTsk(t *testing.T) {
	task := NewTask("test", "test new task", time.Now(), "mohsen")
	if task.Name != "test" {
		t.Errorf("NewTask fail,reason %s", task.Name)
	}
	if task.Description != "test new task" {
		t.Errorf("NewTask fail, reason: %s", "description new task")
	}
	if task.Started {
		t.Errorf("NewTask fail, resson %s", "not started yet")
	}
}
