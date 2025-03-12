package category

import (
	"rest-api-go/internal/module/category/handler"
	"rest-api-go/internal/module/category/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Initialize - Fungsi untuk menginisialisasi modul category
func Initialize(db *gorm.DB, router *gin.RouterGroup) {
	// Initialize service
	categoryService := service.NewCategoryService(db)

	// Initialize handler
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Register routes
	handler.RegisterRoutes(router, categoryHandler)
}
