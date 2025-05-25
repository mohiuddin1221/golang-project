package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohiuddin1221/golang-project/internal/database"
	"github.com/mohiuddin1221/golang-project/internal/routes"
	"github.com/mohiuddin1221/golang-project/pkg/logger"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	database.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// Register user-related routes
	routes.UserRoutes(router)

	// Get the port from environment variables or use the default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Start the server
	if err := router.Run(":" + port); err != nil {
		logger.Log.Fatal("Failed to start server:", err)
	}
}
