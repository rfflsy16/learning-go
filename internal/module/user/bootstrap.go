package user                                  // Mendefinisikan package user

import (
	"rest-api-go/internal/module/user/handler"  // Mengimpor package handler dari modul user
	"rest-api-go/internal/module/user/service"  // Mengimpor package service dari modul user

	"github.com/gin-gonic/gin"                  // Mengimpor framework web Gin
	"gorm.io/gorm"                              // Mengimpor ORM GORM
)

// Initialize - Fungsi untuk menginisialisasi modul user
func Initialize(db *gorm.DB, router *gin.RouterGroup) {  // Fungsi untuk inisialisasi modul dengan parameter database dan router
	// Initialize service
	userService := service.NewUserService(db)    // Membuat instance service user dengan menyuntikkan database

	// Initialize handler
	userHandler := handler.NewUserHandler(userService)  // Membuat instance handler dengan menyuntikkan service

	// Register routes
	handler.RegisterRoutes(router, userHandler)     // Mendaftarkan route untuk modul user
}


// {{{ Penjelasan Fungsi Initialize }}}

/*
## Penjelasan Detail
File bootstrap.go ini berfungsi sebagai titik masuk (entry point) untuk modul user. Berikut penjelasan detailnya:

1. Tujuan : File ini menerapkan pola Dependency Injection untuk menginisialisasi dan menghubungkan komponen-komponen dalam modul user.
2. Alur Kerja :

	- Menerima koneksi database ( db ) dan grup router ( router ) dari aplikasi utama
	- Membuat instance service dengan menyuntikkan database
	- Membuat instance handler dengan menyuntikkan service
	- Mendaftarkan route API untuk modul user
3. Pola Desain :

	- Dependency Injection : Komponen-komponen (service, handler) menerima dependensi mereka dari luar
	- Separation of Concerns : Pemisahan tanggung jawab antara service (logika bisnis) dan handler (penanganan HTTP)
4. Hubungan dengan Aplikasi Utama :

	- Fungsi Initialize dipanggil dari main.go aplikasi utama
	- Menerima koneksi database dan grup router yang sudah diinisialisasi
Struktur ini konsisten dengan modul-modul lain (category, product) yang telah dijelaskan sebelumnya, menunjukkan pendekatan modular yang konsisten dalam arsitektur aplikasi. Setiap modul mengikuti pola yang sama, yang membuat kode lebih mudah dipahami dan dipelihara.
*/