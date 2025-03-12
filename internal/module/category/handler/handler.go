package handler

import (
    "net/http"
    "rest-api-go/internal/module/category/entity"
    "rest-api-go/internal/module/category/service"
    "rest-api-go/pkg/utils"
    "strconv"

    "github.com/gin-gonic/gin"
)

type CategoryHandler struct {
    service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
    return &CategoryHandler{service}
}

func (h *CategoryHandler) Create(c *gin.Context) {
    var category entity.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
        return
    }

    if err := h.service.Create(&category); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, utils.SuccessResponse(category))
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    category, err := h.service.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Category not found"))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(category))
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
    categories, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(categories))
}

func (h *CategoryHandler) Update(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    var category entity.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
        return
    }

    // Set ID dari parameter URL
    category.ID = uint(id)

    if err := h.service.Update(&category); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(category))
}

func (h *CategoryHandler) Delete(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Category deleted successfully"))
}
