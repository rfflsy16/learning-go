package product                                // Mendefinisikan package product

import (
	"rest-api-go/internal/module/product/handler"  // Mengimpor package handler dari modul product
	"rest-api-go/internal/module/product/service"  // Mengimpor package service dari modul product

	"github.com/gin-gonic/gin"                     // Mengimpor framework web Gin
	"gorm.io/gorm"                                 // Mengimpor ORM GORM
)

// Initialize - Fungsi untuk menginisialisasi modul product
func Initialize(db *gorm.DB, router *gin.RouterGroup) {  // Fungsi untuk inisialisasi modul dengan parameter database dan router
	// Initialize service
	productService := service.NewProductService(db)    // Membuat instance service product dengan menyuntikkan database

	// Initialize handler
	productHandler := handler.NewProductHandler(productService)  // Membuat instance handler dengan menyuntikkan service

	// Register routes
	handler.RegisterRoutes(router, productHandler)     // Mendaftarkan route untuk modul product
}


// {{{ Penjelasan Fungsi Initialize }}}

/*
## Penjelasan Detail
File bootstrap.go ini berfungsi sebagai titik masuk (entry point) untuk modul product. Berikut penjelasan detailnya:

1. Tujuan : File ini menerapkan pola Dependency Injection untuk menginisialisasi dan menghubungkan komponen-komponen dalam modul product.
2. Alur Kerja :

	- Menerima koneksi database ( db ) dan grup router ( router ) dari aplikasi utama
	- Membuat instance service dengan menyuntikkan database
	- Membuat instance handler dengan menyuntikkan service
	- Mendaftarkan route API untuk modul product
3. Pola Desain :

	- Dependency Injection : Komponen-komponen (service, handler) menerima dependensi mereka dari luar
	- Separation of Concerns : Pemisahan tanggung jawab antara service (logika bisnis) dan handler (penanganan HTTP)
4. Hubungan dengan Aplikasi Utama :

	- Fungsi Initialize dipanggil dari main.go aplikasi utama
	- Menerima koneksi database dan grup router yang sudah diinisialisasi
Struktur ini mirip dengan modul category yang telah dijelaskan sebelumnya, menunjukkan konsistensi dalam arsitektur aplikasi. Setiap modul (product, category, user) mengikuti pola yang sama, yang membuat kode lebih mudah dipahami dan dipelihara.
*/