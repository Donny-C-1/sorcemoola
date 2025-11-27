package main

import (
	"log"
	"os"

	"github.com/donny-c-1/sorcemoola/server/auth"
	"github.com/donny-c-1/sorcemoola/server/database"
	"github.com/donny-c-1/sorcemoola/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env vars
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning! Could not load .env file: %v", err)
	}

	// Set gin mode based on env vars
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Run database migrations
	if err := database.Migrate(); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	// Initialize gin router
	router := gin.Default()

	// Prevent Untrusted Proxies
	router.SetTrustedProxies([]string{})

	// Setup Cors
	router.Use(auth.CORSMiddleware)

	// Setup routes
	routes.SetupRoutes(router)
	port := os.Getenv("PORT")

	// Get port
	if port == "" {
		log.Fatalf("Port must be set in env")
	}

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
