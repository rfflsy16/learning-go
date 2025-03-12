package handler                                // Mendefinisikan package handler untuk modul product

import (
    "net/http"                                 // Package untuk konstanta HTTP
    "rest-api-go/internal/module/product/entity"  // Mengimpor entity product
    "rest-api-go/internal/module/product/service" // Mengimpor service product
    "rest-api-go/pkg/utils"                    // Mengimpor utilitas aplikasi
    "strconv"                                  // Package untuk konversi string

    "github.com/gin-gonic/gin"                 // Framework web Gin
)

type ProductHandler struct {                   // Mendefinisikan struct handler
    service *service.ProductService            // Dependency service
}

func NewProductHandler(service *service.ProductService) *ProductHandler {  // Constructor untuk handler
    return &ProductHandler{service}            // Mengembalikan instance handler dengan service yang diinjeksi
}

func (h *ProductHandler) Create(c *gin.Context) {  // Handler untuk membuat product baru
    var product entity.Product                 // Variabel untuk menampung data product dari request
    if err := c.ShouldBindJSON(&product); err != nil {  // Binding JSON request ke struct product
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))  // Respons error jika binding gagal
        return
    }

    if err := h.service.Create(&product); err != nil {  // Memanggil service untuk membuat product
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusCreated, utils.SuccessResponse(product))  // Respons sukses dengan data product
}

func (h *ProductHandler) GetByID(c *gin.Context) {  // Handler untuk mendapatkan product berdasarkan ID
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    product, err := h.service.GetByID(uint(id))  // Memanggil service untuk mendapatkan product
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Product not found"))  // Respons error jika tidak ditemukan
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(product))  // Respons sukses dengan data product
}

func (h *ProductHandler) GetAll(c *gin.Context) {  // Handler untuk mendapatkan semua product
    products, err := h.service.GetAll()  // Memanggil service untuk mendapatkan semua product
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(products))  // Respons sukses dengan data products
}

func (h *ProductHandler) Update(c *gin.Context) {  // Handler untuk memperbarui product
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    var product entity.Product  // Variabel untuk menampung data product dari request
    if err := c.ShouldBindJSON(&product); err != nil {  // Binding JSON request ke struct product
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))  // Respons error jika binding gagal
        return
    }

    // Set ID dari parameter URL
    product.ID = uint(id)  // Mengatur ID product dari parameter URL

    if err := h.service.Update(&product); err != nil {  // Memanggil service untuk memperbarui product
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(product))  // Respons sukses dengan data product yang diperbarui
}

func (h *ProductHandler) Delete(c *gin.Context) {  // Handler untuk menghapus product
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {  // Memanggil service untuk menghapus product
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Product deleted successfully"))  // Respons sukses dengan pesan
}

// Add this method to the existing ProductHandler struct

func (h *ProductHandler) GetByCategoryID(c *gin.Context) {  // Handler untuk mendapatkan product berdasarkan CategoryID
    categoryID, err := strconv.ParseUint(c.Param("categoryId"), 10, 32)  // Mengambil dan mengkonversi parameter categoryId
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid Category ID"))  // Respons error jika CategoryID tidak valid
        return
    }

    products, err := h.service.GetByCategoryID(uint(categoryID))  // Memanggil service untuk mendapatkan product berdasarkan CategoryID
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(products))  // Respons sukses dengan data products
}


// {{{ Penjelasan Fungsi RegisterRoutes }}}

/*
## Penjelasan Detail
File handler.go ini berisi implementasi handler HTTP untuk modul Product. Berikut penjelasan detailnya:

1. Tujuan : File ini menangani HTTP request/response untuk operasi CRUD pada entitas Product.
2. Pola Desain :

    - MVC (Model-View-Controller) : Handler bertindak sebagai Controller yang menghubungkan HTTP request dengan logika bisnis.
    - Dependency Injection : Service diinjeksi ke dalam handler melalui constructor.
3. Operasi CRUD :

    - Create : Membuat product baru dari data JSON request
    - GetByID : Mendapatkan product berdasarkan ID dari parameter URL
    - GetAll : Mendapatkan semua product
    - Update : Memperbarui product berdasarkan ID dan data JSON request
    - Delete : Menghapus product berdasarkan ID
    - GetByCategoryID : Mendapatkan product berdasarkan CategoryID (fitur tambahan)
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
Struktur handler ini mirip dengan modul Category yang telah dijelaskan sebelumnya, menunjukkan konsistensi dalam arsitektur aplikasi. Perbedaan utama adalah adanya method tambahan GetByCategoryID yang memungkinkan pencarian product berdasarkan kategori.
*/