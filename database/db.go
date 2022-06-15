// Package database provides database instantiation functions
package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

//createConnection creates a database connection instance
//It returns the pointer to the connection
func createConnection() *sql.DB {
	dsn := "user=postgres dbname=postgres password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

//ConnectWithDB create a singleton instance of database connection
//The function verifies the existence of the instance and if it's not create it calls createConnection()
//to create the connection
//It returns the instance of the database connection
func ConnectWithDB() *sql.DB {
	if db == nil {
		once.Do(func() {
			db = createConnection()
		})
	}
	return db
}
