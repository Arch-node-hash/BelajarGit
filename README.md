# Belajar Git — Todo CLI

Proyek Go kecil untuk belajar Git. Aplikasi ini adalah pengelola tugas
sederhana berbasis terminal, hanya memakai standard library.

## Struktur

```
.
├── main.go                  # CLI: add, list, done, rm
├── internal/todo/
│   ├── task.go              # model Task dan Status
│   ├── store.go             # penyimpanan JSON (tulis atomik)
│   └── store_test.go        # unit test
├── go.mod
└── Makefile
```

## Pemakaian

```bash
go run . add "belajar git"
go run . add "baca dokumentasi Go"
go run . list
go run . done 1
go run . rm 2
```

Data tersimpan di `tasks.json` pada direktori kerja. Berkas ini sudah
diabaikan Git karena data lokal tidak pantas masuk repositori.

## Pengembangan

```bash
make test    # go test ./...
make vet     # go vet ./...
make fmt     # gofmt -w .
```
