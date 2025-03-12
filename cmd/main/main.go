package main // Mendefinisikan package utama untuk aplikasi

import ( // Mengimpor package yang dibutuhkan
	"log"                                  // Package untuk logging
	"rest-api-go/internal/module/category" // Modul category dari aplikasi
	"rest-api-go/internal/module/product"  // Modul product dari aplikasi
	"rest-api-go/internal/module/user"     // Modul user dari aplikasi
	"rest-api-go/pkg/config"               // Package konfigurasi
	"rest-api-go/pkg/database"             // Package database
	"rest-api-go/pkg/middleware"           // Package middleware

	"github.com/gin-gonic/gin" // Framework web Gin
)                                             

func main() {                                 // Fungsi utama yang dijalankan saat program dimulai
	
	// Load config                            
	cfg := config.LoadConfig()                // Memuat konfigurasi aplikasi

	// Connect to database                    
	db := database.Connect()                  // Menghubungkan ke database

	// Setup router                           
	r := gin.Default()                        // Membuat router Gin dengan konfigurasi default
	r.Use(middleware.CORS())                  // Menggunakan middleware CORS

	// API routes                             
	api := r.Group("/api")                    // Membuat grup route dengan prefix "/api"

	// Initialize modules                     
	user.Initialize(db, api)                  // Menginisialisasi modul user
	product.Initialize(db, api)               // Menginisialisasi modul product
	category.Initialize(db, api)              // Menginisialisasi modul category

	// Start server                           
	log.Printf("🚀 Server running on port %s", cfg.ServerPort)  // Menampilkan pesan server berjalan
	r.Run(":" + cfg.ServerPort)               // Menjalankan server pada port yang ditentukan
}




// {{{ Penjelasan Struktur Aplikasi }}}

/*
## Penjelasan Struktur Aplikasi
Aplikasi ini menggunakan arsitektur modular dengan pola MVC (Model-View-Controller) yang dimodifikasi. Berikut penjelasan detailnya:

### 1. Struktur Utama
- Package main : Titik masuk aplikasi yang menginisialisasi semua komponen
- Package internal : Berisi kode yang spesifik untuk aplikasi ini
- Package pkg : Berisi kode yang bisa digunakan kembali di aplikasi lain
### 2. Modul-modul Aplikasi
Setiap modul (user, product, category) memiliki struktur yang sama:

- Entity : Mendefinisikan struktur data dan validasi
- Service : Berisi logika bisnis
- Handler : Menangani HTTP request dan response
- Bootstrap : Menginisialisasi modul dan mendaftarkan route
### 3. Package Pendukung
- Config : Konfigurasi aplikasi
- Database : Koneksi database
- Middleware : Fungsi middleware seperti CORS
- Utils : Fungsi utilitas seperti format response
### Alur Kerja Aplikasi
1. Inisialisasi : main.go memuat konfigurasi dan menghubungkan ke database
2. Setup Router : Membuat router Gin dan menerapkan middleware
3. Registrasi Route : Setiap modul mendaftarkan route-nya sendiri
4. Menjalankan Server : Server HTTP dijalankan pada port yang ditentukan
### Cara Kerja Request
1. Request masuk ke router Gin
2. Middleware diproses (seperti CORS)
3. Request diteruskan ke handler yang sesuai
4. Handler memanggil service untuk logika bisnis
5. Service berinteraksi dengan database melalui entity
6. Response dikembalikan ke client dalam format JSON
### Langkah Selanjutnya untuk Belajar
1. Pelajari Gin Framework : Memahami routing, middleware, dan binding data
2. Pelajari GORM : ORM yang digunakan untuk interaksi database
3. Memahami Struktur Modular : Bagaimana aplikasi dipisahkan menjadi modul-modul
4. Validasi dan Error Handling : Bagaimana aplikasi memvalidasi input dan menangani error
5. Autentikasi dan Otorisasi : Menambahkan sistem login dan kontrol akses
Aplikasi ini adalah contoh yang baik dari REST API dengan Go yang menggunakan praktik terbaik seperti:

- Pemisahan kepentingan (separation of concerns)
- Dependency injection
- Validasi input
- Penanganan error yang konsisten
- Struktur respons yang standar
Untuk pengembangan lebih lanjut, Anda bisa menambahkan fitur seperti autentikasi JWT, caching, logging yang lebih baik, dan pengujian unit.
*/