package model

import (
	"testing"
	"time"
)

func TestTaskDefaults(t *testing.T) {
	now := time.Now()
	task := Task{
		ID:          1,
		Title:       "Sample",
		Description: "Sample Desc",
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if task.ID != 1 {
		t.Errorf("expected ID=1, got %d", task.ID)
	}
	if task.Title != "Sample" {
		t.Errorf("expected Title=Sample, got %s", task.Title)
	}
	if task.Completed {
		t.Errorf("expected Completed=false by default")
	}
}