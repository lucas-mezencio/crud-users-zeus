package main

import (
	"crud_tasks/models"
	"fmt"
)

func main() {
	fmt.Println(models.GetUserById(1))
	//routes.HandleRoutes()
}
