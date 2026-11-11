# Panduan Git dari Nol

Semua perintah di bawah **Anda jalankan sendiri** di terminal, dari folder
`BelajarGit`. Urutannya mengikuti riwayat proyek ini: satu commit untuk satu
perubahan yang utuh.

## 0. Kondisi folder saat ini

Folder ini sudah pernah di-`git init` dan punya satu commit percobaan
(`chore: mulai proyek Go dan atur .gitignore`). Kalau ingin latihan benar-benar
dari nol, hapus dulu repositorinya:

```bash
rm -rf .git
```

Perintah itu hanya menghapus riwayat Git, **bukan** kode Go Anda.

## 1. Atur identitas (sekali saja per komputer)

Nama dan email ini yang akan tercatat di setiap commit:

```bash
git config --global user.name "Nama Anda"
git config --global user.email "email@contoh.com"
git config --global init.defaultBranch main
```

Cek hasilnya:

```bash
git config --global --list
```

## 2. Inisialisasi repositori

```bash
git init -b main
```

Arti: buat repositori Git di folder ini, dengan branch utama bernama `main`.
Folder `.git/` yang muncul adalah database Git. Jangan diedit manual.

## 3. Lihat keadaan

```bash
git status
```

Semua berkas masih **untracked** (belum dikenal Git). Perhatikan bahwa
`tasks.json` nanti tidak akan muncul karena sudah masuk `.gitignore`.

## 4. Commit pertama: fondasi proyek

Pilih berkas yang mau di-commit, lalu commit:

```bash
git add .gitignore go.mod
git status
git commit -m "chore: mulai proyek Go dan atur .gitignore"
```

- `git add` memindahkan berkas ke **staging area** (daftar yang akan di-commit).
- `git commit` menyimpan staging area menjadi satu titik riwayat.

Lihat hasilnya:

```bash
git log --oneline
```

## 5. Commit kedua: lapisan domain

```bash
git add internal/todo/task.go internal/todo/store.go
git diff --staged
git commit -m "feat(todo): tambah model dan penyimpanan tugas"
```

`git diff --staged` menampilkan isi yang **akan** di-commit. Bandingkan dengan
`git diff` yang menampilkan perubahan yang **belum** di-stage.

## 6. Commit ketiga: CLI

```bash
git add main.go
git commit -m "feat: tambah CLI todo"
```

## 7. Commit keempat: test

```bash
git add internal/todo/store_test.go
git commit -m "test(todo): uji operasi penyimpanan"
```

Sekarang riwayatnya:

```bash
git log --oneline --graph --decorate
```

## 8. Latihan branch dan merge

Buat branch untuk dokumentasi, kerjakan di sana, lalu gabungkan:

```bash
git switch -c docs/readme
git add README.md Makefile
git commit -m "docs: tambah README dan Makefile"
git switch main
git merge --no-ff docs/readme
git branch -d docs/readme
```

- `git switch -c` membuat branch baru sekaligus pindah ke sana.
- `--no-ff` memaksa merge commit, supaya jejak branch terlihat di grafik.
- `git branch -d` menghapus branch yang sudah selesai di-merge.

Lihat grafiknya:

```bash
git log --oneline --graph --all
```

## 9. Pastikan `.gitignore` bekerja

```bash
go run . add "coba gitignore"
git status
```

`tasks.json` tidak boleh muncul. Kalau muncul, berarti pola di `.gitignore`
salah. Hapus berkasnya kalau tidak diperlukan:

```bash
rm tasks.json
```

## 10. Perintah harian yang paling sering dipakai

```bash
git status                 # apa yang berubah?
git diff                   # perubahan yang belum di-stage
git add <berkas>           # pilih perubahan
git commit -m "pesan"      # simpan ke riwayat
git log --oneline          # riwayat ringkas
git switch <branch>        # pindah branch
git switch -c <branch>     # buat branch baru
```

Membatalkan perubahan:

```bash
git restore <berkas>            # buang perubahan yang belum di-stage
git restore --staged <berkas>   # keluarkan dari staging, berkas tetap aman
git commit --amend              # perbaiki commit terakhir (pesan/isi)
```

## 11. Push ke GitHub (opsional)

Buat repositori kosong di GitHub, lalu:

```bash
git remote add origin https://github.com/USERNAME/NAMA-REPO.git
git push -u origin main
```

Setelah itu cukup `git push` untuk mengirim commit berikutnya.

## 12. Pola commit yang baik

- Satu commit = satu perubahan logis (jangan campur fitur dan refactor).
- Pesan commit memakai bentuk perintah: `tambah`, `perbaiki`, `ubah`.
- Awalan yang umum: `feat:`, `fix:`, `docs:`, `test:`, `chore:`.
- Commit sebelum pindah tugas, supaya riwayat mudah dibaca dan di-`revert`.
