package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertUser(t *testing.T) {
	userName := "teste 1"
	userEmail := "teste1@mail.com"
	userPassword := "test1-pass"
	userPhone := "34 912341234"
	t.Run("insert user on database", func(t *testing.T) {
		id, _ := InsertUser(userName, userEmail, userPassword, userPhone)
		got := GetUserById(id)
		want := User{id, userName, userEmail, userPassword, userPhone}
		assert.Equal(t, want, got)
		DeleteUserById(id)
	})
	t.Run("insert user with unique field already registered", func(t *testing.T) {
		id, _ := InsertUser(userName, userEmail, userPassword, userPhone)
		_, err := InsertUser(userName, userEmail, userPassword, userPhone)
		assert.Error(t, err, "insert user with unique field already registered")
		DeleteUserById(id)
	})
}

func TestEditUserById(t *testing.T) {
	currentUser := User{
		Name:     "teste321",
		Email:    "teste321@mail",
		Password: "senha321",
		Phone:    "34 22222222",
	}
	modifiedUser := User{
		Name:     "teste123",
		Email:    "teste123@mail",
		Password: "senha123",
		Phone:    "34 phone",
	}

	t.Run("edit user", func(t *testing.T) {
		currentUser.ID, _ = InsertUser(
			currentUser.Name,
			currentUser.Email,
			currentUser.Password,
			currentUser.Phone,
		)
		modifiedUser.ID = currentUser.ID
		EditUser(modifiedUser)
		expectedUser := GetUserById(modifiedUser.ID)
		defer DeleteUserById(modifiedUser.ID)
		assert.Equal(t, expectedUser, modifiedUser)
	})
}
