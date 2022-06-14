package models

import (
	"crud_tasks/database"
	"errors"
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
	err = db.Close()
	if err != nil {
		log.Panicln(err.Error())
	}
	return users
}

func GetUserById(id int) User {
	db := database.ConnectWithDB()
	row := db.QueryRow("select * from zeus.users where id = $1", id)
	var resId int64
	var name, email, password, phone string
	err := row.Scan(&resId, &name, &email, &password, &phone)
	if err != nil { // ErrNoRows
		log.Panicln("User not found", err.Error())
	}
	return User{int(resId), name, email, password, phone}
}

func InsertUser(name, email, password, phone string) (int, error) {
	db := database.ConnectWithDB()
	var lastId int64
	err := db.QueryRow(`insert into zeus.users(nome, email, senha, telefone) values ($1, $2, $3, $4) returning id`,
		name, email,
		password,
		phone).Scan(&lastId)

	if err != nil {
		return 0, errors.New("insert user with unique field already registered")
	}

	err = db.Close()
	if err != nil {
		log.Panicln(err.Error())
	}
	return int(lastId), nil
}
