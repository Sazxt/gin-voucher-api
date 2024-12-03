package database

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "os"
)

func Connect() *sql.DB {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Gagal memuat file .env: %v", err)
    }

    dbDriver := os.Getenv("DB_DRIVER")
    dbSource := os.Getenv("DB_SOURCE")
    if dbDriver == "" || dbSource == "" {
        log.Fatalf("Variabel DB_DRIVER atau DB_SOURCE tidak ditemukan di file .env")
    }
    db, err := sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatalf("Gagal koneksi ke database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Tidak bisa menghubungi database: %v", err)
    }

    return db
}
