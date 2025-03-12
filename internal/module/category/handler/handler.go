package handler                                // Mendefinisikan package handler untuk modul category

import (
    "net/http"                                 // Package untuk konstanta HTTP
    "rest-api-go/internal/module/category/entity"  // Mengimpor entity category
    "rest-api-go/internal/module/category/service" // Mengimpor service category
    "rest-api-go/pkg/utils"                    // Mengimpor utilitas aplikasi
    "strconv"                                  // Package untuk konversi string

    "github.com/gin-gonic/gin"                 // Framework web Gin
)

type CategoryHandler struct {                  // Mendefinisikan struct handler
    service *service.CategoryService           // Dependency service
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {  // Constructor untuk handler
    return &CategoryHandler{service}           // Mengembalikan instance handler dengan service yang diinjeksi
}

func (h *CategoryHandler) Create(c *gin.Context) {  // Handler untuk membuat category baru
    var category entity.Category               // Variabel untuk menampung data category dari request
    if err := c.ShouldBindJSON(&category); err != nil {  // Binding JSON request ke struct category
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))  // Respons error jika binding gagal
        return
    }

    if err := h.service.Create(&category); err != nil {  // Memanggil service untuk membuat category
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusCreated, utils.SuccessResponse(category))  // Respons sukses dengan data category
}

func (h *CategoryHandler) GetByID(c *gin.Context) {  // Handler untuk mendapatkan category berdasarkan ID
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    category, err := h.service.GetByID(uint(id))  // Memanggil service untuk mendapatkan category
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Category not found"))  // Respons error jika tidak ditemukan
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(category))  // Respons sukses dengan data category
}

func (h *CategoryHandler) GetAll(c *gin.Context) {  // Handler untuk mendapatkan semua category
    categories, err := h.service.GetAll()  // Memanggil service untuk mendapatkan semua category
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(categories))  // Respons sukses dengan data categories
}

func (h *CategoryHandler) Update(c *gin.Context) {  // Handler untuk memperbarui category
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    var category entity.Category  // Variabel untuk menampung data category dari request
    if err := c.ShouldBindJSON(&category); err != nil {  // Binding JSON request ke struct category
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))  // Respons error jika binding gagal
        return
    }

    // Set ID dari parameter URL
    category.ID = uint(id)  // Mengatur ID category dari parameter URL

    if err := h.service.Update(&category); err != nil {  // Memanggil service untuk memperbarui category
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(category))  // Respons sukses dengan data category yang diperbarui
}

func (h *CategoryHandler) Delete(c *gin.Context) {  // Handler untuk menghapus category
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {  // Memanggil service untuk menghapus category
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Category deleted successfully"))  // Respons sukses dengan pesan
}

// {{{ Penjelasan Fungsi RegisterRoutes }}}

/*
## Penjelasan Detail
File handler.go ini berisi implementasi handler HTTP untuk modul Category. Berikut penjelasan detailnya:

1. Tujuan : File ini menangani HTTP request/response untuk operasi CRUD pada entitas Category.
2. Pola Desain :

    - MVC (Model-View-Controller) : Handler bertindak sebagai Controller yang menghubungkan HTTP request dengan logika bisnis.
    - Dependency Injection : Service diinjeksi ke dalam handler melalui constructor.
3. Operasi CRUD :

    - Create : Membuat category baru dari data JSON request
    - GetByID : Mendapatkan category berdasarkan ID dari parameter URL
    - GetAll : Mendapatkan semua category
    - Update : Memperbarui category berdasarkan ID dan data JSON request
    - Delete : Menghapus category berdasarkan ID
4. Alur Request :

    - Menerima HTTP request dari router
    - Memvalidasi dan mengekstrak data dari request
    - Memanggil service untuk melakukan operasi bisnis
    - Mengembalikan respons HTTP yang sesuai
5. Penanganan Error :

    - Error binding JSON: Status 400 Bad Request
    - Error validasi atau tidak ditemukan: Status 404 Not Found
    - Error internal: Status 500 Internal Server Error
6. Format Respons :

    - Menggunakan fungsi utils.SuccessResponse dan utils.ErrorResponse untuk format respons yang konsisten
    - Respons sukses berisi data dan status sukses
    - Respons error berisi pesan error dan status gagal
Handler ini mengimplementasikan prinsip "thin controller" di mana handler hanya bertanggung jawab untuk menangani HTTP request/response, sementara logika bisnis sepenuhnya berada di service.
*/