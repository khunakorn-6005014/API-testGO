// backend/handler/task.go
package handler

import (
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/khunakorn-6005014/API-testGO/backend/model"
  "github.com/khunakorn-6005014/API-testGO/backend/store"
)

// TaskHandler holds a store.TaskStore so handlers can call CRUD methods.
type TaskHandler struct {
  Store store.TaskStore
}

// NewTaskHandler returns a TaskHandler wired to the given store.
func NewTaskHandler(s store.TaskStore) *TaskHandler {
  return &TaskHandler{Store: s}
}

// List returns all tasks.
func (h *TaskHandler) List(c *gin.Context) {
  tasks, err := h.Store.List()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, tasks)
}

// Get returns one task by ID.
func (h *TaskHandler) Get(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
    return
  }

  task, err := h.Store.Get(id)
  if err != nil {
    if err == store.ErrNotFound {
      c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    } else {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }
    return
  }

  c.JSON(http.StatusOK, task)
}

// Create adds a new task.
func (h *TaskHandler) Create(c *gin.Context) {
  var input model.Task
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  task, err := h.Store.Create(input)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, task)
}

// Update modifies an existing task.
func (h *TaskHandler) Update(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
    return
  }

  var input model.Task
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  task, err := h.Store.Update(id, input)
  if err != nil {
    if err == store.ErrNotFound {
      c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    } else {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }
    return
  }

  c.JSON(http.StatusOK, task)
}

// Delete removes a task.
func (h *TaskHandler) Delete(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
    return
  }

  if err := h.Store.Delete(id); err != nil {
    if err == store.ErrNotFound {
      c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    } else {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }
    return
  }

  c.Status(http.StatusNoContent)
}