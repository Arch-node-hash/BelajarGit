package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	// ErrNotFound dikembalikan saat ID tugas tidak ada.
	ErrNotFound = errors.New("tugas tidak ditemukan")
	// ErrEmptyTitle dikembalikan saat judul tugas kosong.
	ErrEmptyTitle = errors.New("judul tugas tidak boleh kosong")
)

// Store menyimpan daftar tugas pada sebuah berkas JSON.
//
// Perubahan (Add, Complete, Remove) hanya mengubah data di memori.
// Panggil Save untuk menuliskannya ke berkas.
type Store struct {
	path  string
	tasks []Task
}

// NewStore membuat Store yang membaca dan menulis ke path.
func NewStore(path string) *Store {
	return &Store{path: path, tasks: []Task{}}
}

// Path mengembalikan lokasi berkas penyimpanan.
func (s *Store) Path() string {
	return s.path
}

// Load membaca tugas dari berkas. Berkas yang belum ada dianggap kosong.
func (s *Store) Load() error {
	data, err := os.ReadFile(s.path)
	if errors.Is(err, os.ErrNotExist) {
		s.tasks = []Task{}
		return nil
	}
	if err != nil {
		return fmt.Errorf("membaca %s: %w", s.path, err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return fmt.Errorf("mengurai %s: %w", s.path, err)
	}
	if tasks == nil {
		tasks = []Task{}
	}
	s.tasks = tasks
	return nil
}

// Save menulis seluruh tugas ke berkas secara atomik, sehingga berkas
// tidak pernah tertinggal dalam keadaan setengah tertulis.
func (s *Store) Save() error {
	data, err := json.MarshalIndent(s.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("menyusun JSON: %w", err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(s.path)
	tmp, err := os.CreateTemp(dir, ".tasks-*.json")
	if err != nil {
		return fmt.Errorf("membuat berkas sementara: %w", err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName) // tidak berpengaruh setelah rename berhasil

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("menulis berkas sementara: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("menutup berkas sementara: %w", err)
	}
	if err := os.Rename(tmpName, s.path); err != nil {
		return fmt.Errorf("mengganti %s: %w", s.path, err)
	}
	return nil
}

// All mengembalikan salinan seluruh tugas.
func (s *Store) All() []Task {
	out := make([]Task, len(s.tasks))
	copy(out, s.tasks)
	return out
}

// Add menambahkan tugas baru dan mengembalikannya.
func (s *Store) Add(title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	task := Task{
		ID:        s.nextID(),
		Title:     title,
		Status:    StatusTodo,
		CreatedAt: time.Now().UTC(),
	}
	s.tasks = append(s.tasks, task)
	return task, nil
}

// Complete menandai tugas dengan id sebagai selesai.
func (s *Store) Complete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Status = StatusDone
			return nil
		}
	}
	return fmt.Errorf("tugas %d: %w", id, ErrNotFound)
}

// Remove menghapus tugas dengan id.
func (s *Store) Remove(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("tugas %d: %w", id, ErrNotFound)
}

// nextID mengembalikan ID terkecil yang belum dipakai.
func (s *Store) nextID() int {
	max := 0
	for _, task := range s.tasks {
		if task.ID > max {
			max = task.ID
		}
	}
	return max + 1
}
