package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"tasktrek/config"
	"tasktrek/routes"
)

func main() {
	// Connect DB
	config.ConnectDB()

	// Set Gin mode depending on environment
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Register routes
	routes.UserRoutes(r)
	routes.TaskRoutes(r)

	// Get port from Railway or default to 8080 locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("🚀 Listening on %s\n", addr)

	// Start server
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
