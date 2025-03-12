package product

import (
	"rest-api-go/internal/module/product/handler"
	"rest-api-go/internal/module/product/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Initialize - Fungsi untuk menginisialisasi modul product
func Initialize(db *gorm.DB, router *gin.RouterGroup) {
	// Initialize service
	productService := service.NewProductService(db)

	// Initialize handler
	productHandler := handler.NewProductHandler(productService)

	// Register routes
	handler.RegisterRoutes(router, productHandler)
}
