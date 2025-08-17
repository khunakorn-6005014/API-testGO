package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khunakorn-6005014/API-testGO/backend/model"
	"github.com/khunakorn-6005014/API-testGO/backend/store"
)

func TestListTasks(t *testing.T) {
	s := store.NewMemoryStore()
	s.Create(model.Task{Title: "Test Task", Description: "desc"})
	h := NewTaskHandler(s)

	r := gin.Default()
	r.GET("/tasks", h.List)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var tasks []model.Task
	if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestCreateTask(t *testing.T) {
	s := store.NewMemoryStore()
	h := NewTaskHandler(s)

	r := gin.Default()
	r.POST("/tasks", h.Create)

	body := []byte(`{"title":"New Task","description":"something"}`)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}