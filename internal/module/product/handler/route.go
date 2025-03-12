package handler                                // Mendefinisikan package handler untuk modul product

import (
    "github.com/gin-gonic/gin"                 // Mengimpor framework web Gin
)

func RegisterRoutes(router *gin.RouterGroup, handler *ProductHandler) {  // Fungsi untuk mendaftarkan route
    products := router.Group("/products")      // Membuat grup route dengan prefix "/products"
    {
        products.POST("", handler.Create)      // Mendaftarkan endpoint POST untuk membuat product baru
        products.GET("/:id", handler.GetByID)  // Mendaftarkan endpoint GET dengan parameter id untuk mendapatkan product berdasarkan ID
        products.GET("", handler.GetAll)       // Mendaftarkan endpoint GET untuk mendapatkan semua product
        products.GET("/category/:categoryId", handler.GetByCategoryID)  // Mendaftarkan endpoint GET untuk mendapatkan product berdasarkan kategori
        products.PUT("/:id", handler.Update)   // Mendaftarkan endpoint PUT dengan parameter id untuk memperbarui product
        products.DELETE("/:id", handler.Delete)  // Mendaftarkan endpoint DELETE dengan parameter id untuk menghapus product
    }
}


// {{{ Penjelasan Fungsi RegisterRoutes }}}

/*
## Penjelasan Detail
File route.go ini berisi konfigurasi routing untuk modul Product. Berikut penjelasan detailnya:

1. Tujuan : File ini mendaftarkan endpoint API untuk operasi CRUD pada entitas Product.
2. Struktur Routing :

    - Semua endpoint dikelompokkan di bawah prefix /products
    - Endpoint dikelompokkan berdasarkan metode HTTP (POST, GET, PUT, DELETE)
3. Endpoint API :

    - POST /products : Membuat product baru
    - GET /products/:id : Mendapatkan product berdasarkan ID
    - GET /products : Mendapatkan semua product
    - GET /products/category/:categoryId : Mendapatkan product berdasarkan kategori
    - PUT /products/:id : Memperbarui product berdasarkan ID
    - DELETE /products/:id : Menghapus product berdasarkan ID
4. Parameter URL :

    - :id : Parameter dinamis untuk ID product
    - :categoryId : Parameter dinamis untuk ID kategori
5. Handler Mapping :

    - Setiap endpoint dipetakan ke method handler yang sesuai
    - Handler menerima dan memproses request HTTP
6. Fitur Tambahan :

    - Endpoint /products/category/:categoryId memungkinkan pencarian product berdasarkan kategori
    - Ini menunjukkan bagaimana routing dapat digunakan untuk mengimplementasikan relasi antar entitas
Pendekatan ini memisahkan konfigurasi routing dari implementasi handler, yang membuat kode lebih terorganisir dan mudah dipelihara. Ini juga memungkinkan pengujian routing secara terpisah dari implementasi handler.
*/