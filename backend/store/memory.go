// backend/store/memory.go
package store

import (
  "sync"
  "time"

  "github.com/khunakorn-6005014/API-testGO/backend/model"
)

// MemoryStore holds tasks in memory and satisfies TaskStore.
type MemoryStore struct {
  mu     sync.Mutex
  lastID int
  data   map[int]*model.Task
}

// NewMemoryStore returns an empty in-memory store.
func NewMemoryStore() *MemoryStore {
  return &MemoryStore{
    data: make(map[int]*model.Task),
  }
}

func (s *MemoryStore) List() ([]model.Task, error) {
  s.mu.Lock()
  defer s.mu.Unlock()

  tasks := make([]model.Task, 0, len(s.data))
  for _, t := range s.data {
    tasks = append(tasks, *t)
  }
  return tasks, nil
}

func (s *MemoryStore) Get(id int) (model.Task, error) {
  s.mu.Lock()
  defer s.mu.Unlock()

  t, ok := s.data[id]
  if !ok {
    return model.Task{}, ErrNotFound
  }
  return *t, nil
}

func (s *MemoryStore) Create(t model.Task) (model.Task, error) {
  s.mu.Lock()
  defer s.mu.Unlock()

  s.lastID++
  t.ID = s.lastID
  t.CreatedAt = time.Now()
  t.UpdatedAt = t.CreatedAt
  s.data[t.ID] = &t
  return t, nil
}

func (s *MemoryStore) Update(id int, t model.Task) (model.Task, error) {
  s.mu.Lock()
  defer s.mu.Unlock()

  existing, ok := s.data[id]
  if !ok {
    return model.Task{}, ErrNotFound
  }
  existing.Title = t.Title
  existing.Description = t.Description
  existing.Completed = t.Completed
  existing.UpdatedAt = time.Now()
  return *existing, nil
}

func (s *MemoryStore) Delete(id int) error {
  s.mu.Lock()
  defer s.mu.Unlock()

  if _, ok := s.data[id]; !ok {
    return ErrNotFound
  }
  delete(s.data, id)
  return nil
}