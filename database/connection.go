package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // required by sql.Open for postgres
)

func ConnectDB(DATABASE_URL string) *sql.DB {
	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to make a connection to database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Ping to database failed:", err)
	}
	fmt.Println("✅ Success connect to database")
	return db
}
