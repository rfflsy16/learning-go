package seed                                  // Mendefinisikan package seed

import (
    "encoding/json"                           // Package untuk encoding/decoding JSON
    "fmt"                                     // Package untuk formatting dan output
    "log"                                     // Package untuk logging
    "os"                                      // Package untuk operasi sistem
    "rest-api-go/internal/module/user/entity"  // Mengimpor entity user

    "gorm.io/gorm"                            // Mengimpor ORM GORM
)

// Users - fungsi untuk seed data user
func Users(db *gorm.DB) {                     // Fungsi untuk seed data user dengan parameter database
    // Drop table if exists
    err := db.Migrator().DropTable(&entity.User{})  // Menghapus tabel user jika ada
    if err != nil {
        log.Fatal("Error dropping table:", err)  // Log error dan hentikan program jika gagal
    }
    fmt.Println("🗑️  Old user tables dropped successfully")  // Pesan sukses menghapus tabel

    // Auto migrate
    err = db.AutoMigrate(&entity.User{})      // Membuat tabel user berdasarkan struct entity
    if err != nil {
        log.Fatal("Error migrating user table:", err)  // Log error dan hentikan program jika gagal
    }
    fmt.Println("📝 New user table created successfully")  // Pesan sukses membuat tabel

    // Read JSON file
    data, err := os.ReadFile("data/users.json")  // Membaca file JSON data user
    if err != nil {
        log.Fatal("Error reading users.json:", err)  // Log error dan hentikan program jika gagal
    }

    // Parse JSON data
    var users []entity.User                   // Variabel untuk menampung data user dari JSON
    err = json.Unmarshal(data, &users)        // Mengkonversi JSON ke slice struct User
    if err != nil {
        log.Fatal("Error parsing users.json:", err)  // Log error dan hentikan program jika gagal
    }

    // Seed data
    for _, user := range users {              // Iterasi setiap user dari data JSON
        if err := db.Create(&user).Error; err != nil {  // Menyimpan user ke database
            log.Fatal("Error seeding user:", err)  // Log error dan hentikan program jika gagal
        }
    }

    fmt.Println("🌱 User data seeded successfully!")  // Pesan sukses seed data
}



// {{{ Penjelasan Fungsi Seeder }}}

/*
## Penjelasan Detail
File user.go ini berisi fungsi untuk melakukan seed data user ke dalam database. Berikut penjelasan detailnya:

1. Tujuan : File ini digunakan untuk mengisi database dengan data awal (seed data) untuk entitas User.
2. Alur Kerja :

    - Menghapus tabel User yang sudah ada (jika ada)
    - Membuat tabel baru berdasarkan struktur entity User
    - Membaca data dari file JSON
    - Mengkonversi data JSON ke slice struct User
    - Menyimpan data User ke database
3. Fitur Database :

    - DropTable : Menghapus tabel yang sudah ada
    - AutoMigrate : Membuat tabel berdasarkan struktur struct
    - Create : Menyimpan data ke database
4. Penanganan File :

    - Membaca file JSON dari path data/users.json
    - Menggunakan json.Unmarshal untuk mengkonversi JSON ke struct Go
5. Penanganan Error :

    - Menggunakan log.Fatal untuk menghentikan program jika terjadi error
    - Menampilkan pesan error yang informatif
6. Feedback :

    - Menampilkan pesan sukses dengan emoji untuk setiap tahap proses
    - Memudahkan pengguna untuk memahami progress seed data
Fungsi ini mengikuti pola yang sama dengan fungsi seed untuk Category dan Product, menunjukkan konsistensi dalam pendekatan seeding data di seluruh aplikasi.
*/