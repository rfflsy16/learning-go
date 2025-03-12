package handler

import (
    "net/http"
    "rest-api-go/internal/module/user/entity"
    "rest-api-go/internal/module/user/service"
    "rest-api-go/pkg/utils"
    "strconv"

    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
    return &UserHandler{service}
}

func (h *UserHandler) Create(c *gin.Context) {
    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
        return
    }

    if err := h.service.Create(&user); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, utils.SuccessResponse(user))
}

func (h *UserHandler) GetByID(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    user, err := h.service.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(user))
}

func (h *UserHandler) GetAll(c *gin.Context) {
    users, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(users))
}

func (h *UserHandler) Update(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
        return
    }

    // Set ID dari parameter URL
    user.ID = uint(id)

    if err := h.service.Update(&user); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(user))
}

func (h *UserHandler) Delete(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("User deleted successfully"))
}
