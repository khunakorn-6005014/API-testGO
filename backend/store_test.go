package store

import (
  "testing"
)

func TestMemoryStore_CRUD(t *testing.T) {
  s := NewMemoryStore()

  // Create
  t1, _ := s.Create(...)
  if t1.ID != 1 { t.Errorf("expected ID 1, got %d", t1.ID) }

  // Get
  if _, err := s.Get(1); err != nil {
    t.Fatalf("unexpected error: %v", err)
  }

  // Update & Delete assertions...
}