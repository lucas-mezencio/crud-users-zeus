package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

//fazer singleton
func createConnection() *sql.DB {
	dsn := "user=postgres dbname=postgres password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func ConnectWithDB() *sql.DB {
	if db == nil {
		once.Do(func() {
			db = createConnection()
		})
	}
	return db
}
