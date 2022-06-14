package handlers

import (
	"crud_tasks/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetUsers())
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot understand this id",
		})
		return
	}
	user := models.GetUserById(id)
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "not found",
		})
		return
	}
	c.JSON(http.StatusOK, models.GetUserById(id))
}

func DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot understand this id",
		})
		return
	}
	user := models.GetUserById(id)
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "user not found to delete",
		})
		return
	}

	models.DeleteUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"status": "user deleted successfully",
		"data":   user,
	})
}
