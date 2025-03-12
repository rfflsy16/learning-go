package category // Mendefinisikan package category

import (
	"rest-api-go/internal/module/category/handler" // Mengimpor package handler dari modul category
	"rest-api-go/internal/module/category/service" // Mengimpor package service dari modul category

	"github.com/gin-gonic/gin" // Mengimpor framework web Gin
	"gorm.io/gorm"             // Mengimpor ORM GORM
)

// Initialize - Fungsi untuk menginisialisasi modul category
func Initialize(db *gorm.DB, router *gin.RouterGroup) {  // Fungsi untuk inisialisasi modul dengan parameter database dan router
	// Initialize service
	categoryService := service.NewCategoryService(db)    // Membuat instance service category dengan menyuntikkan database

	// Initialize handler
	categoryHandler := handler.NewCategoryHandler(categoryService)  // Membuat instance handler dengan menyuntikkan service

	// Register routes
	handler.RegisterRoutes(router, categoryHandler)     // Mendaftarkan route untuk modul category
}


// {{{ Penjelasan Fungsi Initialize }}}

/*
## Penjelasan Detail
File bootstrap.go ini berfungsi sebagai titik masuk (entry point) untuk modul category. Berikut penjelasan detailnya:

1. Tujuan : File ini menerapkan pola Dependency Injection untuk menginisialisasi dan menghubungkan komponen-komponen dalam modul category.
2. Alur Kerja :

	- Menerima koneksi database ( db ) dan grup router ( router ) dari aplikasi utama
	- Membuat instance service dengan menyuntikkan database
	- Membuat instance handler dengan menyuntikkan service
	- Mendaftarkan route API untuk modul category
3. Pola Desain :

	- Dependency Injection : Komponen-komponen (service, handler) menerima dependensi mereka dari luar
	- Separation of Concerns : Pemisahan tanggung jawab antara service (logika bisnis) dan handler (penanganan HTTP)
4. Hubungan dengan Aplikasi Utama :

	- Fungsi Initialize dipanggil dari main.go aplikasi utama
	- Menerima koneksi database dan grup router yang sudah diinisialisasi
## Konsep Penting
1. Modularitas : Setiap modul (category, product, user) memiliki struktur yang sama dan dapat diinisialisasi secara independen.
2. Dependency Injection : Dependensi (database, service) disuntikkan dari luar, bukan dibuat di dalam komponen, yang memudahkan pengujian dan penggantian implementasi.
3. Layering : Aplikasi mengikuti pola arsitektur berlapis:

	- Handler (Presentation Layer): Menangani HTTP request/response
	- Service (Business Logic Layer): Berisi logika bisnis
	- Entity (Data Access Layer): Berinteraksi dengan database
Pola desain ini membuat kode lebih terstruktur, mudah diuji, dan mudah dipelihara karena setiap komponen memiliki tanggung jawab yang jelas dan terisolasi.
*/