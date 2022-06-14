package routes

import (
	"crud_tasks/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/users", handlers.GetAllUsers)

	err := r.Run()
	if err != nil {
		log.Panicln(err.Error())
	}
}
