// backend/main.go
package main

import (
  
  "github.com/gin-gonic/gin"
  "github.com/khunakorn-6005014/API-testGO/backend/handler"
  "github.com/khunakorn-6005014/API-testGO/backend/store"
)

func main() {
 // 1. Initialize Gin router
  r := gin.Default()
  // 2. Create & inject the TaskStore
  ts := handler.NewTaskHandler(store.NewMemoryStore())
  // 3. Define routes
  r.GET("/tasks", ts.List)
  r.GET("/tasks/:id", ts.Get)
  r.POST("/tasks", ts.Create)
  r.PUT("/tasks/:id", ts.Update)
  r.DELETE("/tasks/:id", ts.Delete)
  // 4. Run on :8080
  r.Run(":8080")
}

