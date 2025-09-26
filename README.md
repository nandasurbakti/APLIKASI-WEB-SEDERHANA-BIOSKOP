# APLIKASI-WEB-SEDERHANA-BIOSKOP

Project ini adalah REST API sederhana untuk mengelola data bioskop. Dibuat dengan bahasa Go menggunakan Gin Framework dan PostgreSQL sebagai database.

Fitur :
Autentikasi JWT untuk proteksi endpoint.
CRUD data bioskop:
Create Bioskop
Read semua bioskop
Read bioskop berdasarkan ID
Update bioskop
Delete bioskop
Migration database menggunakan sql-migrate.

Struktur Folder:
bioskop-app/
│
├─ config/ # konfigurasi DB dan env
├─ controllers/ # handler request
├─ database/ # migration setup
├─ models/ # struct & query DB
├─ routers/ # definisi routing
├─ sql_migrations/ # file migration SQL
├─ main.go
└─ README.md

PS: Maaf berantakan masih pemula
