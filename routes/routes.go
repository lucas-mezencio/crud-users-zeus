//Package routes is responsible for routing the application
package routes

import (
	"crud_tasks/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

// HandleRoutes defines the routes of the application using the gin framework
func HandleRoutes() {
	r := gin.Default()

	r.GET("/users", handlers.GetAllUsers)
	r.GET("/users/:id", handlers.GetUserById)
	r.DELETE("/users/:id", handlers.DeleteUserById)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.EditUserById)
	r.PATCH("/users/:id", handlers.PatchUserById)
	r.NoRoute(handlers.NoRoute)

	err := r.Run()
	if err != nil {
		log.Panicln(err.Error())
	}
}
