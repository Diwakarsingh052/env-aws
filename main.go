package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in the environment variables")
	}

	// Initialize the Gin router
	r := gin.Default()

	// Define the ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong v1",
		})
	})

	// Start the server on the specified port
	log.Printf("Starting server on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
