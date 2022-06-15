package handlers

import (
	"crud_tasks/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"err": err.Error(),
			},
		)
		return
	}

	id, err := models.InsertUser(user.Name, user.Email, user.Password, user.Phone)
	user.ID = id
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}

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

func EditUserById(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	userCheck := models.GetUserById(user.ID)
	if userCheck == (models.User{}) {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	models.EditUser(user)
	c.JSON(http.StatusOK, user)
}

func PatchUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot understand this id",
		})
		return
	}

	currentUser := models.GetUserById(id)
	if currentUser == (models.User{}) {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var user models.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	user.ID = currentUser.ID
	if user.Name == "" {
		user.Name = currentUser.Name
	}

	if user.Email == "" {
		user.Email = currentUser.Email
	}

	if user.Password == "" {
		user.Password = currentUser.Password
	}

	if user.Phone == "" {
		user.Phone = currentUser.Phone
	}

	c.JSON(http.StatusOK, user)
}
