package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"belajargit/internal/todo"
)

const storeFile = "tasks.json"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		usage()
		return nil
	}

	store := todo.NewStore(storeFile)
	if err := store.Load(); err != nil {
		return err
	}

	switch args[0] {
	case "add":
		return add(store, args[1:])
	case "list":
		return list(store)
	case "done":
		return complete(store, args[1:])
	case "rm":
		return remove(store, args[1:])
	case "help", "-h", "--help":
		usage()
		return nil
	default:
		usage()
		return fmt.Errorf("perintah tidak dikenal: %q", args[0])
	}
}

func add(store *todo.Store, args []string) error {
	task, err := store.Add(strings.Join(args, " "))
	if err != nil {
		return err
	}
	if err := store.Save(); err != nil {
		return err
	}

	fmt.Printf("ditambahkan #%d: %s\n", task.ID, task.Title)
	return nil
}

func list(store *todo.Store) error {
	tasks := store.All()
	if len(tasks) == 0 {
		fmt.Println("belum ada tugas")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tSTATUS\tJUDUL")
	for _, task := range tasks {
		mark := "[ ]"
		if task.Done() {
			mark = "[x]"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, mark, task.Title)
	}
	return w.Flush()
}

func complete(store *todo.Store, args []string) error {
	id, err := parseID(args)
	if err != nil {
		return err
	}
	if err := store.Complete(id); err != nil {
		return err
	}
	if err := store.Save(); err != nil {
		return err
	}

	fmt.Printf("selesai #%d\n", id)
	return nil
}

func remove(store *todo.Store, args []string) error {
	id, err := parseID(args)
	if err != nil {
		return err
	}
	if err := store.Remove(id); err != nil {
		return err
	}
	if err := store.Save(); err != nil {
		return err
	}

	fmt.Printf("dihapus #%d\n", id)
	return nil
}

func parseID(args []string) (int, error) {
	if len(args) != 1 {
		return 0, errors.New("butuh tepat satu ID tugas")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, fmt.Errorf("ID tidak valid: %q", args[0])
	}
	return id, nil
}

func usage() {
	fmt.Print(`pemakaian:
  todo add <judul>   tambah tugas baru
  todo list          tampilkan semua tugas
  todo done <id>     tandai tugas selesai
  todo rm <id>       hapus tugas
`)
}
