package database

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB(DATABSE_URL string) *sql.DB {
	db, err := sql.Open("postgres", DATABSE_URL)
	if err != nil {
		log.Fatal("Failed to make a connection to database", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Ping to database is failed", err)
	}
	fmt.Println("Success connect to database")
	return db
}
