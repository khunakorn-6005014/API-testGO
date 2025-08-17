package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khunakorn-6005014/API-testGO/backend/handler"
	"github.com/khunakorn-6005014/API-testGO/backend/store"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	ts := handler.NewTaskHandler(store.NewMemoryStore())
	r.GET("/tasks", ts.List)
	r.GET("/tasks/:id", ts.Get)
	r.POST("/tasks", ts.Create)
	r.PUT("/tasks/:id", ts.Update)
	r.DELETE("/tasks/:id", ts.Delete)
	return r
}

func TestListTasks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", w.Code)
	}
}