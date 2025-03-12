package main

import (
	"log"
	"rest-api-go/internal/module/product"
	"rest-api-go/internal/module/user"
	"rest-api-go/pkg/config"
	"rest-api-go/pkg/database"
	"rest-api-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db := database.Connect()

	// Setup router
	r := gin.Default()
	r.Use(middleware.CORS())

	// API routes
	api := r.Group("/api")

	// Initialize modules
	product.Initialize(db, api)
	user.Initialize(db, api)

	// Start server
	log.Printf("🚀 Server running on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
