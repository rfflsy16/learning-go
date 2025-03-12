package handler

import (
    "net/http"
    "rest-api-go/internal/module/product/entity"
    "rest-api-go/internal/module/product/service"
    "rest-api-go/pkg/utils"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ProductHandler struct {
    service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
    return &ProductHandler{service}
}

func (h *ProductHandler) Create(c *gin.Context) {
    var product entity.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
        return
    }

    if err := h.service.Create(&product); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, utils.SuccessResponse(product))
}

func (h *ProductHandler) GetByID(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    product, err := h.service.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Product not found"))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(product))
}

func (h *ProductHandler) GetAll(c *gin.Context) {
    products, err := h.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(products))
}

func (h *ProductHandler) Update(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    var product entity.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
        return
    }

    // Set ID dari parameter URL
    product.ID = uint(id)

    if err := h.service.Update(&product); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(product))
}

func (h *ProductHandler) Delete(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Product deleted successfully"))
}

// Add this method to the existing ProductHandler struct

func (h *ProductHandler) GetByCategoryID(c *gin.Context) {
    categoryID, err := strconv.ParseUint(c.Param("categoryId"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Category ID"))
        return
    }

    products, err := h.service.GetByCategoryID(uint(categoryID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(products))
}
