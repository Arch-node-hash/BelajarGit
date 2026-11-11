// Package todo menyediakan model tugas dan penyimpanan berbasis JSON.
package todo

import "time"

// Status merepresentasikan keadaan sebuah tugas.
type Status string

const (
	// StatusTodo berarti tugas belum dikerjakan.
	StatusTodo Status = "todo"
	// StatusDone berarti tugas sudah selesai.
	StatusDone Status = "done"
)

// Task adalah satu tugas dalam daftar.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Done melaporkan apakah tugas sudah selesai.
func (t Task) Done() bool {
	return t.Status == StatusDone
}
