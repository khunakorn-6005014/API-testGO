// backend/store/store.go
package store

import (
    "errors"

    "github.com/khunakorn-6005014/API-testGO/backend/model"
)
// ErrNotFound is returned when a Task can’t be found.
var ErrNotFound = errors.New("task not found")

// TaskStore describes the CRUD interface for tasks.
type TaskStore interface {
  List() ([]model.Task, error)
  Get(id int) (model.Task, error)
  Create(t model.Task) (model.Task, error)
  Update(id int, t model.Task) (model.Task, error)
  Delete(id int) error
}