package main

import (
	"tasktrek/config"
	"tasktrek/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()
	routes.UserRoutes(router)
	routes.TaskRoutes(router)

	router.Run(":8080") // Start server on localhost:8080
}
