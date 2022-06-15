// Package handlers provides handlers (controllers) for each model on the application
package handlers

import (
	"crud_tasks/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateUser handles the route to create users:
//		POST /users
// If the request body can't be read returns a Bad Request status
// If the user already exists returns also a Bad Request status
// Returns the user created as json on body and a Created status
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

// GetAllUsers handles the route to get all users on db:
//		GET /users
// Returns a list of users as json, could be an empty list and Ok status
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetUsers())
}

// GetUserById handles the route to get a user by its id:
// 		GET /users/id
// If the request body can't be read returns a Bad Request status
// If the user doesn't exist returns a Not Found status
// Returns a user as json and Ok status
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

// DeleteUserById handles the route to delete a user by its id:
//		DELETE /users/id
// If the request body can't be read returns a Bad Request status
// If the user doesn't exist returns a Not Found status
// Returns a json body with a success message and the deleted user as data and Ok status
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

// EditUserById handles the route to edit a user by its id:
//		PUT /users/id
// If the request body can't be read returns a Bad Request status
// If the user doesn't exist returns a Not Found status
// Returns the update user data as json and Ok status
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

// PatchUserById handles the route to edit a user via patch method by its id:
//		PATCH /users/id
// If the request body can't be read returns a Bad Request status
// If the user doesn't exist returns a Not Found status
// If the id can't be read returns a Bad Request status
// Returns the updated user with an Ok status
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

// NoRoute handles any not specified route on application
// Returns a body with a not found message and a Not Found status
func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": "Page not Found",
	})
}
