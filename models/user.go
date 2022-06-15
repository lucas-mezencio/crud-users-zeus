// Package models includes the models of the application and its database operations
package models

import (
	"crud_tasks/database"
	"database/sql"
	"errors"
	"log"
)

// User defines the user model for the application purpose and database table users
type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// GetUsers consults the database to search for all users.
// It returns a list (slice) of User.
func GetUsers() []User {
	db := database.ConnectWithDB()
	users := make([]User, 0)

	rows, err := db.Query("select * from zeus.users order by id asc")
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

// GetUserById search on the database to find a specific user given its id as argument.
// If the user don't exists the function returns an empty one
// Returns the user with the given id.
func GetUserById(id int) User {
	db := database.ConnectWithDB()
	row := db.QueryRow("select * from zeus.users where id = $1", id)
	var resId int64
	var name, email, password, phone string
	err := row.Scan(&resId, &name, &email, &password, &phone)
	if err == sql.ErrNoRows { // ErrNoRows
		return User{}
	} else if err != nil {
		log.Panicln("User not found", err.Error())

	}
	return User{int(resId), name, email, password, phone}
}

// InsertUser inserts a user on database by giving its attributes.
// If exists a user with same email or same phone in database it returns 0 and an error.
// Returns the id of the new user.
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

// DeleteUserById deletes an user on database given its id
func DeleteUserById(id int) {
	db := database.ConnectWithDB()
	defer db.Close()

	_, err := db.Exec("delete from zeus.users where id = $1", id)
	if err != nil {
		log.Panicln(err.Error())
	}
}

// EditUser is responsible for update user data on database given a new struct User with the
//matching id existent on database.
func EditUser(user User) {
	db := database.ConnectWithDB()
	defer db.Close()

	_, err := db.Exec(
		"update zeus.users set nome = $1, email = $2, senha = $3, telefone = $4 where id = $5",
		user.Name,
		user.Email,
		user.Password,
		user.Phone,
		user.ID,
	)
	if err != nil {
		log.Panicf("Can't edit this user - Error: %q\n", err.Error())
	}
}
