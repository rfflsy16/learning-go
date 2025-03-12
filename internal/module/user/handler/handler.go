package handler                                // Mendefinisikan package handler untuk modul user

import (
    "net/http"                                 // Package untuk konstanta HTTP
    "rest-api-go/internal/module/user/entity"  // Mengimpor entity user
    "rest-api-go/internal/module/user/service" // Mengimpor service user
    "rest-api-go/pkg/utils"                    // Mengimpor utilitas aplikasi
    "strconv"                                  // Package untuk konversi string

    "github.com/gin-gonic/gin"                 // Framework web Gin
)

type UserHandler struct {                      // Mendefinisikan struct handler
    service *service.UserService               // Dependency service
}

func NewUserHandler(service *service.UserService) *UserHandler {  // Constructor untuk handler
    return &UserHandler{service}               // Mengembalikan instance handler dengan service yang diinjeksi
}

func (h *UserHandler) Create(c *gin.Context) {  // Handler untuk membuat user baru
    var user entity.User                       // Variabel untuk menampung data user dari request
    if err := c.ShouldBindJSON(&user); err != nil {  // Binding JSON request ke struct user
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))  // Respons error jika binding gagal
        return
    }

    if err := h.service.Create(&user); err != nil {  // Memanggil service untuk membuat user
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusCreated, utils.SuccessResponse(user))  // Respons sukses dengan data user
}

func (h *UserHandler) GetByID(c *gin.Context) {  // Handler untuk mendapatkan user berdasarkan ID
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    user, err := h.service.GetByID(uint(id))  // Memanggil service untuk mendapatkan user
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))  // Respons error jika tidak ditemukan
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(user))  // Respons sukses dengan data user
}

func (h *UserHandler) GetAll(c *gin.Context) {  // Handler untuk mendapatkan semua user
    users, err := h.service.GetAll()  // Memanggil service untuk mendapatkan semua user
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(users))  // Respons sukses dengan data users
}

func (h *UserHandler) Update(c *gin.Context) {  // Handler untuk memperbarui user
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    var user entity.User  // Variabel untuk menampung data user dari request
    if err := c.ShouldBindJSON(&user); err != nil {  // Binding JSON request ke struct user
        c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))  // Respons error jika binding gagal
        return
    }

    // Set ID dari parameter URL
    user.ID = uint(id)  // Mengatur ID user dari parameter URL

    if err := h.service.Update(&user); err != nil {  // Memanggil service untuk memperbarui user
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse(user))  // Respons sukses dengan data user yang diperbarui
}

func (h *UserHandler) Delete(c *gin.Context) {  // Handler untuk menghapus user
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)  // Mengambil dan mengkonversi parameter ID
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))  // Respons error jika ID tidak valid
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {  // Memanggil service untuk menghapus user
        c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))  // Respons error jika gagal
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("User deleted successfully"))  // Respons sukses dengan pesan
}


// {{{ Penjelasan Fungsi RegisterRoutes }}}

/*
## Penjelasan Detail
File handler.go ini berisi implementasi handler HTTP untuk modul User. Berikut penjelasan detailnya:

1. Tujuan : File ini menangani HTTP request/response untuk operasi CRUD pada entitas User.
2. Pola Desain :

    - MVC (Model-View-Controller) : Handler bertindak sebagai Controller yang menghubungkan HTTP request dengan logika bisnis.
    - Dependency Injection : Service diinjeksi ke dalam handler melalui constructor.
3. Operasi CRUD :

    - Create : Membuat user baru dari data JSON request
    - GetByID : Mendapatkan user berdasarkan ID dari parameter URL
    - GetAll : Mendapatkan semua user
    - Update : Memperbarui user berdasarkan ID dan data JSON request
    - Delete : Menghapus user berdasarkan ID
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
Struktur handler ini konsisten dengan modul-modul lain (category, product) yang telah dijelaskan sebelumnya, menunjukkan pendekatan yang konsisten dalam arsitektur aplikasi.
*/