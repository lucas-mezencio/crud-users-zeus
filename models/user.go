package models

import (
	"crud_tasks/database"
	"log"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Phone    string
}

func GetUsers() []User {
	db := database.ConnectWithDB()
	var users []User

	rows, err := db.Query("select * from zeus.users asc")
	if err != nil {
		log.Panicln(err.Error())
	}

	for rows.Next() {
		var id int64
		var name, email, password, phone string
		if err := rows.Scan(&id, &name, &email, &password, &phone); err != nil {
			log.Panicln(err.Error())
		}
		users = append(users, User{int(id), name, email, password, phone})
	}
	err = rows.Close()
	if err != nil {
		log.Panicln(err.Error())
	}
	err = db.Close()
	if err != nil {
		log.Panicln(err.Error())
	}
	return users
}

//func GetUserById(id int) User {
//
//}

func InsertUser(name, email, password, phone string) int {
	db := database.ConnectWithDB()

	result, err := db.Exec(`insert into users(nome, email, senha, telefone) values ($1, $2, $3, $4)`, name, email,
		password,
		phone)

	if err != nil {
		log.Panicln(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
	}

	err = db.Close()
	if err != nil {
		log.Panicln(err.Error())
	}
	return int(lastId)
}
