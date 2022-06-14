package routes

import (
	"crud_tasks/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
