package routes

import (
	"crud_tasks/handlers"
	"crud_tasks/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func setupTestingRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func TestGetAllUsers(t *testing.T) {
	r := setupTestingRoutes()

	t.Run("test get all users status code", func(t *testing.T) {
		r.GET("/users", handlers.GetAllUsers)

		req, _ := http.NewRequest("GET", "/users", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})
}

func TestGetUserById(t *testing.T) {
	r := setupTestingRoutes()

	t.Run("test get user by id", func(t *testing.T) {

		id, _ := models.InsertUser("test get by id", "getbyid@mail.com", "pass", "a number")
		userExpect := models.User{
			ID:       id,
			Name:     "test get by id",
			Email:    "getbyid@mail.com",
			Password: "pass",
			Phone:    "a number",
		}
		defer models.DeleteUserById(id)

		r.GET("/users/:id", handlers.GetUserById)

		req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		var user models.User
		_ = json.Unmarshal(res.Body.Bytes(), &user)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, userExpect.Email, user.Email)
	})

	t.Run("test get user by id where id not on db", func(t *testing.T) {
		r.GET("/users/:id", handlers.GetUserById)

		req, _ := http.NewRequest("GET", "/users/0", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		expect := `{"status":"not found"}`
		resBody, _ := ioutil.ReadAll(res.Body)
		got := string(resBody)
		assert.Equal(t, expect, got, "erro response expected")
	})
}