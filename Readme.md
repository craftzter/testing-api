<center>

# Dokumentasi penggunaan api

</center>

## 1. langkah langkah dalam penggunaan api
- Clone github repository
```
https://github.com/craftzter/testing-api
```
- Kemudian jalankan `go mod tidy`, `docker-compose up -d` kode pertama untuk mengunduh depedensi dari golang code kedua untuk pull database dan menjadikan project golang ini menjadi docker,
- Selanjutnya jalankan migration kedalam database sebelum itu buat dulu copy `.example.env menjadi .env` kemudian jalankan migration dari folder migrations/ atau mengunakan `make up` untuk generate schema database

```
Ringkasan URL Akses

http://localhost:8080/users/register → register user

http://localhost:8080/users/login → login & dapatkan JWT

http://localhost:8080/users/{id} → get user by ID

http://localhost:8080/users/list → list semua user

http://localhost:8080/users/profile/{id} → update profile (JWT required)
register memerlukan 3 field yaitu username, email dan password,
login --> email dan password, untuk update profile menggunakan token jwt,
catatan: profile itu masih bisa update id orang lain karena masih implementasi simple
```
