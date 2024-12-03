# Gin Voucher API
Gin Voucher API adalah aplikasi backend sederhana yang menyediakan layanan untuk mengelola brand, voucher, dan transaksi penukaran voucher. Dibangun menggunakan Go dengan framework Gin Gonic.

1. Clone Repository
git clone https://github.com/Sazxt/gin-voucher-api.git
cd gin-voucher-api

2. Instal Dependensi
go mod tidy

3. Konfigurasi Database
```
DB_DRIVER=postgres   # atau mysql
DB_SOURCE="host=localhost user=username password=password dbname=dbname port=5432 sslmode=disable TimeZone=Asia/Jakarta"
```

4. Migrasi Database
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -path ./migrations -database "${DB_SOURCE}" up

Atau lebih gampang langsung aja tempel sql tersebut di query Navicat , saya melakukan ini salin dan tempel di sql query lalu run.

5. Jalankan Aplikasi

go run main.go


Menjalankan Unit Test

go test ./test/... -v


.
├── handlers          # Logic untuk request API
├── models            # Struktur data untuk database
├── migrations        # File SQL migrasi
├── test              # Unit test
├── main.go           # Entry point aplikasi
├── go.mod            # File dependensi
└── README.md         # Dokumentasi
