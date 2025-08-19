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

	// Gin mode
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Routes
	routes.UserRoutes(r)
	routes.TaskRoutes(r)

	// PORT (Railway sets PORT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("🚀 Listening on %s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
