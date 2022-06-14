package handlers

import (
	"crud_tasks/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetUsers())
}
