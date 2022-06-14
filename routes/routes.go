package routes

import (
	"crud_tasks/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/users", handlers.GetAllUsers)
	r.GET("/users/:id", handlers.GetUserById)
	r.DELETE("/users/:id", handlers.DeleteUserById)

	err := r.Run()
	if err != nil {
		log.Panicln(err.Error())
	}
}
