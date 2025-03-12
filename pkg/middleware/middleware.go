package middleware                            // Mendefinisikan package middleware

import (
    "github.com/gin-gonic/gin"                // Mengimpor framework web Gin
)

func Logger() gin.HandlerFunc {               // Fungsi untuk middleware logging
    return gin.Logger()                       // Mengembalikan middleware logger bawaan Gin
}

func CORS() gin.HandlerFunc {                 // Fungsi untuk middleware CORS (Cross-Origin Resource Sharing)
    return func(c *gin.Context) {             // Mengembalikan fungsi handler middleware
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")  // Mengizinkan akses dari semua origin
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")  // Mengizinkan metode HTTP tertentu
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")  // Mengizinkan header tertentu

        if c.Request.Method == "OPTIONS" {    // Jika request adalah OPTIONS (preflight request)
            c.AbortWithStatus(204)            // Mengembalikan status 204 (No Content) dan menghentikan chain middleware
            return
        }

        c.Next()                              // Melanjutkan ke middleware atau handler berikutnya
    }
}



// {{{ Penjelasan Fungsi CORS }}}

/*
## Penjelasan Detail
File middleware.go ini berisi implementasi middleware untuk aplikasi web Gin. Berikut penjelasan detailnya:

1. Tujuan : File ini menyediakan middleware yang dapat digunakan dalam aplikasi untuk:

    - Logging request HTTP
    - Menangani CORS (Cross-Origin Resource Sharing)
2. Middleware Logger :

    - Menggunakan middleware logger bawaan Gin
    - Mencatat informasi request seperti metode HTTP, path, status code, dan waktu respons
    - Membantu dalam debugging dan monitoring aplikasi
3. Middleware CORS :

    - Menambahkan header CORS ke respons HTTP
    - Mengizinkan akses dari semua origin ( * )
    - Mengizinkan metode HTTP: POST, GET, OPTIONS, PUT, DELETE
    - Mengizinkan header tertentu seperti Content-Type, Authorization, dll.
    - Menangani preflight request (OPTIONS) dengan respons 204 No Content
4. Penanganan Preflight Request :

    - Browser mengirim request OPTIONS sebelum request sebenarnya untuk memeriksa apakah request diizinkan
    - Middleware mengembalikan status 204 (No Content) untuk request OPTIONS
    - Ini memungkinkan browser melanjutkan dengan request sebenarnya
5. Penggunaan :

    - Middleware ini biasanya didaftarkan di main.go atau router utama aplikasi
    - Middleware dijalankan untuk setiap request yang masuk sebelum mencapai handler
Middleware CORS ini sangat penting untuk aplikasi web modern, terutama jika frontend dan backend berjalan di domain atau port yang berbeda.
*/