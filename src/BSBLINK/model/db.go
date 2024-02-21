package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var host, port, user, password, dbname string

	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")

	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}

	log.Println("Successfully connected to database")
	return db
}
