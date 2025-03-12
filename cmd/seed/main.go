package main                        // Mendefinisikan package utama untuk aplikasi seeder

import (
	"log"                       // Package untuk logging
	"rest-api-go/internal/seed" // Mengimpor package seed yang berisi fungsi-fungsi seeding
	"rest-api-go/pkg/database"  // Mengimpor package database untuk koneksi
)

func main() {                       // Fungsi utama yang dijalankan saat program seeder dimulai
	// Connect to database
	db := database.Connect()    // Menghubungkan ke database menggunakan fungsi Connect()
	
	// Seed data (termasuk migrasi)
	// Seed categories first, then products to maintain foreign key integrity
	seed.Categories(db)         // Menjalankan fungsi seeding untuk kategori (termasuk migrasi tabel)
	seed.Products(db)           // Menjalankan fungsi seeding untuk produk (termasuk migrasi tabel)
	seed.Users(db)              // Menjalankan fungsi seeding untuk pengguna (termasuk migrasi tabel)
	
	log.Println("✅ All data migrated and seeded successfully") // Menampilkan pesan sukses setelah semua data berhasil di-seed
}



// {{{ Penjelasan Fungsi Seeder }}}

/*
## Penjelasan Fungsi Seeder
File ini adalah program terpisah yang berfungsi untuk mengisi database dengan data awal (seeding). Berikut penjelasan detailnya:

1. Tujuan : Program ini digunakan untuk mengisi database dengan data awal yang diperlukan aplikasi, seperti kategori, produk, dan pengguna.
2. Alur Kerja :

	- Menghubungkan ke database
	- Menjalankan fungsi seeding untuk kategori terlebih dahulu
	- Kemudian menjalankan fungsi seeding untuk produk (karena produk memiliki foreign key ke kategori)
	- Terakhir menjalankan fungsi seeding untuk pengguna
3. Urutan Seeding : Urutan ini penting karena adanya relasi foreign key. Kategori harus dibuat terlebih dahulu sebelum produk, karena produk mereferensikan kategori.
4. Migrasi Tabel : Setiap fungsi seeding juga melakukan migrasi tabel (membuat atau memperbarui struktur tabel) sebelum mengisi data.
## Cara Menjalankan Seeder
Untuk menjalankan program seeder ini, Anda dapat menggunakan perintah:

```bash
go run cmd/seed/main.go
```

Atau jika Anda telah membangun aplikasi:

```bash
./seed
```

## Langkah Selanjutnya
Setelah memahami seeder, Anda dapat:

1. Mempelajari Struktur Data : Lihat file-file di internal/seed/ untuk memahami bagaimana data distruktur dan dimasukkan ke database.
2. Menambahkan Data Baru : Anda dapat menambahkan data baru ke file JSON yang digunakan oleh seeder.
3. Membuat Seeder Baru : Jika Anda menambahkan entitas baru ke aplikasi, Anda perlu membuat fungsi seeding baru.
4. Memahami Migrasi : Pelajari bagaimana GORM melakukan migrasi tabel secara otomatis berdasarkan struktur struct.
Seeder ini adalah bagian penting dari siklus pengembangan, terutama untuk pengujian dan pengembangan awal, karena memungkinkan Anda untuk dengan cepat mengisi database dengan data yang konsisten.
*/