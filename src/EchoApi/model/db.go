// db.go
package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	var host, port, user, password, dbname string

	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("Successfully connected to the database")

	return DB
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}

func GetUserByEmail(email string) (*User, error) {
	// Query user information based on email
	row := DB.QueryRow("SELECT email, password FROM omni.user WHERE email = $1", email)

	var user User
	err := row.Scan(&user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // User not found
		}
		return nil, err // Other error
	}

	return &user, nil
}
