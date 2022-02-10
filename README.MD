# Golang REST API

Aplikasi ini dibuat untuk proses tahap seleksi recuirment PT. Klik Digital

## Usage ⚠️

### Database 🐁

pertama tama buat database bernama 'golang-jwt' pada MySQL RDBMS

kemudian pada database tersebut, import file sql yang ada pada repository ini.

```bash
golang-jwt.sql
```

### Usage 🧨
Setelah itu, jika database local anda memiliki konfigurasi default ( username = root , password = '', host = 'localhost') maka anda tinggal menjalankan binary build file yang pada repository dan aplikasi pun berjalan.



Namun jika konfigurasi database anda berbeda, atur variable yang ada di file config/database-config.go sesuai dengan konfigurasi anda 

```bash
	dbUser :=  UsernameDBAnda
	dbPass :=  PasswordDBAnda
	dbHost :=  HostDBAnda
	dbName :=  "golang-jwt"
```

setelah itu simpan semua perubahan dan jalankan perintah

```bash
go mod tidy 
```
untuk menginstall semua dependecy yang diperlukan 


selanjutnya jalankan 

```bash
go build *.go
```
untuk mengcompile aplikasi menjadi binary build file 

Terakhir, jalankan aplikasi dengan membuka file binary build "server". atau

jalankan perintah

```bash
go run *.go
```
untuk meng-serve aplikasi dengan terminal

Dokumentasi endpoint ada pada file json Golang JWT REST API.postman_collection.json, lakukan import collection pada aplikasi Postman untuk menampilkan semua dokumentasi.
