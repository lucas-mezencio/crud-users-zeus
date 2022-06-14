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
	})
	t.Run("insert user with unique field already registered", func(t *testing.T) {
		_, err := InsertUser(userName, userEmail, userPassword, userPhone)
		assert.Error(t, err, "insert user with unique field already registered")
	})
}
