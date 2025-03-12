package handler                                // Mendefinisikan package handler untuk modul category

import (
    "github.com/gin-gonic/gin"                 // Mengimpor framework web Gin
)

func RegisterRoutes(router *gin.RouterGroup, handler *CategoryHandler) {  // Fungsi untuk mendaftarkan route
    categories := router.Group("/categories")  // Membuat grup route dengan prefix "/categories"
    {
        categories.POST("", handler.Create)    // Mendaftarkan endpoint POST untuk membuat category baru
        categories.GET("/:id", handler.GetByID)  // Mendaftarkan endpoint GET dengan parameter id untuk mendapatkan category berdasarkan ID
        categories.GET("", handler.GetAll)     // Mendaftarkan endpoint GET untuk mendapatkan semua category
        categories.PUT("/:id", handler.Update)   // Mendaftarkan endpoint PUT dengan parameter id untuk memperbarui category
        categories.DELETE("/:id", handler.Delete)  // Mendaftarkan endpoint DELETE dengan parameter id untuk menghapus category
    }
}


// {{{ Penjelasan Fungsi RegisterRoutes }}}

/*
## Penjelasan Detail
File route.go ini berisi konfigurasi routing untuk modul Category. Berikut penjelasan detailnya:

1. Tujuan : File ini mendaftarkan endpoint API untuk operasi CRUD pada entitas Category.
2. Struktur Routing :

    - Semua endpoint dikelompokkan di bawah prefix /categories
    - Endpoint dikelompokkan berdasarkan metode HTTP (POST, GET, PUT, DELETE)
3. Endpoint API :

    - POST /categories : Membuat category baru
    - GET /categories/:id : Mendapatkan category berdasarkan ID
    - GET /categories : Mendapatkan semua category
    - PUT /categories/:id : Memperbarui category berdasarkan ID
    - DELETE /categories/:id : Menghapus category berdasarkan ID
4. Parameter URL :

    - :id : Parameter dinamis untuk ID category
5. Handler Mapping :

    - Setiap endpoint dipetakan ke method handler yang sesuai
    - Handler menerima dan memproses request HTTP
6. Dependency Injection :

    - Router dan handler diinjeksi dari luar melalui parameter fungsi
Pendekatan ini memisahkan konfigurasi routing dari implementasi handler, yang membuat kode lebih terorganisir dan mudah dipelihara. Ini juga memungkinkan pengujian routing secara terpisah dari implementasi handler.

Dalam arsitektur aplikasi secara keseluruhan, file ini berperan sebagai penghubung antara router utama aplikasi dan handler spesifik untuk modul Category.
*/