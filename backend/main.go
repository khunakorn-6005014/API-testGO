package main

import (
  "github.com/gin-gonic/gin"
  "github.com/khunakorn-6005014/API-testGO/backend/handler"
  "github.com/khunakorn-6005014/API-testGO/backend/store"
)

func main() {
  r := gin.Default()
  ts := handler.NewTaskHandler(store.NewMemoryStore())

  r.GET("/tasks", ts.List)
  r.GET("/tasks/:id", ts.Get)
  r.POST("/tasks", ts.Create)
  r.PUT("/tasks/:id", ts.Update)
  r.DELETE("/tasks/:id", ts.Delete)

  r.Run(":8080")
}