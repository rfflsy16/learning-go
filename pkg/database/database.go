package database                              // Mendefinisikan package database

import (
    "fmt"                                     // Package untuk formatting string
    "log"                                     // Package untuk logging
    "rest-api-go/pkg/config"                  // Mengimpor package config aplikasi

    "gorm.io/driver/mysql"                    // Driver MySQL/MariaDB untuk GORM
    "gorm.io/gorm"                            // ORM GORM
)

func Connect() *gorm.DB {                     // Fungsi untuk membuat koneksi database
    // Load config
    cfg := config.LoadConfig()                // Memuat konfigurasi aplikasi

    // Connect to MariaDB
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",  // Membuat string koneksi (DSN)
        cfg.DBUser,                           // Username database dari konfigurasi
        cfg.DBPassword,                       // Password database dari konfigurasi
        cfg.DBHost,                           // Host database dari konfigurasi
        cfg.DBPort,                           // Port database dari konfigurasi
        cfg.DBName,                           // Nama database dari konfigurasi
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  // Membuka koneksi database dengan GORM
    if err != nil {
        log.Fatal("Failed to connect to database:", err)   // Log error dan hentikan program jika koneksi gagal
    }

    return db                                 // Mengembalikan koneksi database
}



// {{{ Penjelasan Fungsi Connect }}}

/*
## Penjelasan Detail
File database.go ini berisi fungsi untuk membuat koneksi ke database MariaDB/MySQL. Berikut penjelasan detailnya:

1. Tujuan : File ini menyediakan fungsi untuk menginisialisasi dan mengembalikan koneksi database yang digunakan oleh aplikasi.
2. Alur Kerja :

    - Memuat konfigurasi database dari package config
    - Membuat string koneksi (DSN - Data Source Name) dengan format yang sesuai untuk MySQL/MariaDB
    - Membuka koneksi database menggunakan GORM dengan driver MySQL
    - Mengembalikan koneksi database jika berhasil atau menghentikan program jika gagal
3. String Koneksi (DSN) :

    - Format: username:password@tcp(host:port)/dbname?params
    - Parameter tambahan:
    - charset=utf8mb4 : Menggunakan encoding UTF-8 yang mendukung karakter 4-byte
    -  parseTime=True : Mengkonversi kolom TIME/DATE menjadi time.Time Go
    - loc=Local : Menggunakan zona waktu lokal untuk waktu
4. Penanganan Error :

    - Menggunakan log.Fatal untuk menghentikan program jika koneksi database gagal
    - Ini memastikan aplikasi tidak berjalan tanpa koneksi database yang valid
5. Penggunaan :

    - Fungsi ini biasanya dipanggil di main.go untuk menginisialisasi koneksi database
    - Koneksi database kemudian digunakan oleh service dan repository untuk operasi database
Pendekatan ini memisahkan konfigurasi dan inisialisasi database dari kode aplikasi utama, yang membuat kode lebih terorganisir dan mudah dipelihara.
*/