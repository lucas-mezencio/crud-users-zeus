package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	err := r.Run()
	if err != nil {
		log.Panicln(err.Error())
	}
}
