package routes

import (
	"crud_tasks/handlers"
	"crud_tasks/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
		r.GET("/users/id", handlers.GetUserById)

		id, _ := models.InsertUser("test get by id", "getbyid@mail.com", "pass", "a number")
		url := "/users" + strconv.Itoa(id)
		req, _ := http.NewRequest("GET", url, nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})
	//t.Run("test get user by id where id not on db")
}
