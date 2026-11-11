package todo

import (
	"errors"
	"path/filepath"
	"testing"
)

func newStore(t *testing.T) *Store {
	t.Helper()

	store := NewStore(filepath.Join(t.TempDir(), "tasks.json"))
	if err := store.Load(); err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	return store
}

func TestAddThenReload(t *testing.T) {
	store := newStore(t)

	task, err := store.Add("belajar git")
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if task.ID != 1 {
		t.Errorf("ID = %d, ingin 1", task.ID)
	}
	if task.Status != StatusTodo {
		t.Errorf("Status = %q, ingin %q", task.Status, StatusTodo)
	}
	if err := store.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	reloaded := NewStore(store.Path())
	if err := reloaded.Load(); err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	tasks := reloaded.All()
	if len(tasks) != 1 {
		t.Fatalf("len(tasks) = %d, ingin 1", len(tasks))
	}
	if tasks[0].Title != "belajar git" {
		t.Errorf("Title = %q, ingin %q", tasks[0].Title, "belajar git")
	}
}

func TestAddRejectsEmptyTitle(t *testing.T) {
	store := newStore(t)

	if _, err := store.Add("   "); !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("Add() error = %v, ingin ErrEmptyTitle", err)
	}
}

func TestComplete(t *testing.T) {
	store := newStore(t)

	task, err := store.Add("selesai")
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if err := store.Complete(task.ID); err != nil {
		t.Fatalf("Complete() error = %v", err)
	}
	if !store.All()[0].Done() {
		t.Error("tugas seharusnya selesai")
	}

	if err := store.Complete(99); !errors.Is(err, ErrNotFound) {
		t.Errorf("Complete(99) error = %v, ingin ErrNotFound", err)
	}
}

func TestRemove(t *testing.T) {
	store := newStore(t)

	task, err := store.Add("hapus aku")
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	if err := store.Remove(task.ID); err != nil {
		t.Fatalf("Remove() error = %v", err)
	}
	if len(store.All()) != 0 {
		t.Errorf("len(tasks) = %d, ingin 0", len(store.All()))
	}

	if err := store.Remove(task.ID); !errors.Is(err, ErrNotFound) {
		t.Errorf("Remove() error = %v, ingin ErrNotFound", err)
	}
}
