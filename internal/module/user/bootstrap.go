package user

import (
	"rest-api-go/internal/module/user/handler"
	"rest-api-go/internal/module/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Initialize - Fungsi untuk menginisialisasi modul user
func Initialize(db *gorm.DB, router *gin.RouterGroup) {
	// Initialize service
	userService := service.NewUserService(db)

	// Initialize handler
	userHandler := handler.NewUserHandler(userService)

	// Register routes
	handler.RegisterRoutes(router, userHandler)
}
