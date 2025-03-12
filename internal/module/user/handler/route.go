package handler                                // Mendefinisikan package handler untuk modul user

import (
    "github.com/gin-gonic/gin"                 // Mengimpor framework web Gin
)

func RegisterRoutes(router *gin.RouterGroup, handler *UserHandler) {  // Fungsi untuk mendaftarkan route
    users := router.Group("/users")            // Membuat grup route dengan prefix "/users"
    {
        users.POST("", handler.Create)         // Mendaftarkan endpoint POST untuk membuat user baru
        users.GET("/:id", handler.GetByID)     // Mendaftarkan endpoint GET dengan parameter id untuk mendapatkan user berdasarkan ID
        users.GET("", handler.GetAll)          // Mendaftarkan endpoint GET untuk mendapatkan semua user
        users.PUT("/:id", handler.Update)      // Mendaftarkan endpoint PUT dengan parameter id untuk memperbarui user
        users.DELETE("/:id", handler.Delete)   // Mendaftarkan endpoint DELETE dengan parameter id untuk menghapus user
    }
}


// {{{ Penjelasan Fungsi RegisterRoutes }}}

/*
## Penjelasan Detail
File route.go ini berisi konfigurasi routing untuk modul User. Berikut penjelasan detailnya:

1. Tujuan : File ini mendaftarkan endpoint API untuk operasi CRUD pada entitas User.
2. Struktur Routing :

    - Semua endpoint dikelompokkan di bawah prefix /users
    - Endpoint dikelompokkan berdasarkan metode HTTP (POST, GET, PUT, DELETE)
3. Endpoint API :

    - POST /users : Membuat user baru
    - GET /users/:id : Mendapatkan user berdasarkan ID
    - GET /users : Mendapatkan semua user
    - PUT /users/:id : Memperbarui user berdasarkan ID
    - DELETE /users/:id : Menghapus user berdasarkan ID
4. Parameter URL :

    - :id : Parameter dinamis untuk ID user
5. Handler Mapping :

    - Setiap endpoint dipetakan ke method handler yang sesuai
    - Handler menerima dan memproses request HTTP
6. Dependency Injection :

    - Router dan handler diinjeksi dari luar melalui parameter fungsi
Pendekatan ini memisahkan konfigurasi routing dari implementasi handler, yang membuat kode lebih terorganisir dan mudah dipelihara. Ini juga memungkinkan pengujian routing secara terpisah dari implementasi handler.
*/