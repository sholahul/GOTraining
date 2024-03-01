// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Mendapatkan konfigurasi database dari environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Membuat string koneksi database
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Membuat koneksi ke database
	var err error
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Membuat tabel 'users' jika belum ada
	createTable()

	// Menjalankan server web di port 8080
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Menampilkan pesan jika berhasil terkoneksi ke database
	fmt.Fprintln(w, "Aplikasi Golang berhasil terkoneksi ke PostgreSQL!")
}

func createTable() {
	// Perintah SQL untuk membuat tabel 'users'
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);
	`

	// Eksekusi perintah SQL untuk membuat tabel
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
