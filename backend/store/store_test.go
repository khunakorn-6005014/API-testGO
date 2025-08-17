// backend/store/store_test.go
package store

import (
  "testing"

  "github.com/khunakorn-6005014/API-testGO/backend/model"
)

func TestMemoryStore_CRUD(t *testing.T) {
  s := NewMemoryStore()

  // 1) Create—use a real Task literal, not “…”
  input := model.Task{Title: "Test", Description: "desc"}
  created, err := s.Create(input)
  if err != nil {
    t.Fatalf("Create failed: %v", err)
  }
  if created.ID != 1 {
    t.Errorf("expected ID=1, got %d", created.ID)
  }

  // 2) Get
  got, err := s.Get(created.ID)
  if err != nil {
    t.Fatalf("Get failed: %v", err)
  }
  if got.Title != input.Title {
    t.Errorf("expected Title=%q, got %q", input.Title, got.Title)
  }

  // 3) Update
  upd := model.Task{Title: "Updated", Description: "new", Completed: true}
  updated, err := s.Update(created.ID, upd)
  if err != nil {
    t.Fatalf("Update failed: %v", err)
  }
  if updated.Title != "Updated" || !updated.Completed {
    t.Errorf("unexpected updated: %+v", updated)
  }

  // 4) List
  list, _ := s.List()
  if len(list) != 1 {
    t.Errorf("expected 1 task, got %d", len(list))
  }

  // 5) Delete
  if err := s.Delete(created.ID); err != nil {
    t.Fatalf("Delete failed: %v", err)
  }
  if _, err := s.Get(created.ID); err != ErrNotFound {
    t.Errorf("expected ErrNotFound after delete, got %v", err)
  }
}
