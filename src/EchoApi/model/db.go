// db.go
package model

import (
	"database/sql"
	"errors"
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

	var u User
	err := row.Scan(&u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // User not found
		}
		return nil, err // Other error
	}

	return &u, nil
}

func GetAll() ([]User, error) {
	// Ensure the DB connection is properly established
	if DB == nil {
		return nil, errors.New("database connection is nil")
	}

	// Query user information based on email
	rows, err := DB.Query("SELECT id, username, email, role FROM omni.user")
	if err != nil {
		log.Println("ERORRRR")
		return nil, err // Return the error if the query fails
	}
	defer rows.Close()
	log.Println("LEWAT DLU")
	var Data_users []User

	// Iterate over the rows returned by the query
	for rows.Next() {
		var u User // Use the correct struct name
		// Scan each row into the User struct
		err := rows.Scan(&u.Id, &u.Username, &u.Email, &u.Role)
		if err != nil {
			log.Println("ERORRRR")
			return nil, err // Return the error if scanning fails
		}
		Data_users = append(Data_users, u)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Data_users, nil
}
