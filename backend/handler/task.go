package handler

import (
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/khunakorn-6005014/API-testGO/backend/model"
  "github.com/khunakorn-6005014/API-testGO/backend/store"
)

type TaskHandler struct {
  Store store.TaskStore
}

func NewTaskHandler(s store.TaskStore) *TaskHandler {
  return &TaskHandler{Store: s}
}

func (h *TaskHandler) List(c *gin.Context) {
  tasks, _ := h.Store.List()
  c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) Get(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
    return
  }
  t, err := h.Store.Get(id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, t)
}

func (h *TaskHandler) Create(c *gin.Context) {
  var input model.Task
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  task, _ := h.Store.Create(input)
  c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Update(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  var input model.Task
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  task, err := h.Store.Update(id, input)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  if err := h.Store.Delete(id); err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }
  c.Status(http.StatusNoContent)
}