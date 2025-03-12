package handler

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *ProductHandler) {
    products := router.Group("/products")
    {
        products.POST("", handler.Create)
        products.GET("/:id", handler.GetByID)
        products.GET("", handler.GetAll)
        products.GET("/category/:categoryId", handler.GetByCategoryID)
        products.PUT("/:id", handler.Update)
        products.DELETE("/:id", handler.Delete)
    }
}
