package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//fazer singleton
func ConnectWithDB() *sql.DB {
	dsn := "user=postgres dbname=postgres password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
