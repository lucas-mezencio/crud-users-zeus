//Package main is responsible for start the application
package main

import (
	"crud_tasks/routes"
)

//main start the application
func main() {
	routes.HandleRoutes()
}
