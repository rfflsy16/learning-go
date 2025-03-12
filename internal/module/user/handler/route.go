package handler

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *UserHandler) {
    users := router.Group("/users")
    {
        users.POST("", handler.Create)
        users.GET("/:id", handler.GetByID)
        users.GET("", handler.GetAll)
        users.PUT("/:id", handler.Update)
        users.DELETE("/:id", handler.Delete)
    }
}
