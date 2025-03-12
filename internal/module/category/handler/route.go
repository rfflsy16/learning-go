package handler

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *CategoryHandler) {
    categories := router.Group("/categories")
    {
        categories.POST("", handler.Create)
        categories.GET("/:id", handler.GetByID)
        categories.GET("", handler.GetAll)
        categories.PUT("/:id", handler.Update)
        categories.DELETE("/:id", handler.Delete)
    }
}
