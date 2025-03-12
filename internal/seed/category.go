package seed                                  // Mendefinisikan package seed

import (
    "encoding/json"                           // Package untuk encoding/decoding JSON
    "fmt"                                     // Package untuk formatting dan output
    "log"                                     // Package untuk logging
    "os"                                      // Package untuk operasi sistem
    "rest-api-go/internal/module/category/entity"  // Mengimpor entity category

    "gorm.io/gorm"                            // Mengimpor ORM GORM
)

// Categories - fungsi untuk seed data category
func Categories(db *gorm.DB) {                // Fungsi untuk seed data category dengan parameter database
    // Drop table if exists
    err := db.Migrator().DropTable(&entity.Category{})  // Menghapus tabel category jika ada
    if err != nil {
        log.Fatal("Error dropping table:", err)  // Log error dan hentikan program jika gagal
    }
    fmt.Println("🗑️  Old category tables dropped successfully")  // Pesan sukses menghapus tabel

    // Auto migrate
    err = db.AutoMigrate(&entity.Category{})  // Membuat tabel category berdasarkan struct entity
    if err != nil {
        log.Fatal("Error migrating category table:", err)  // Log error dan hentikan program jika gagal
    }
    fmt.Println("📝 New category table created successfully")  // Pesan sukses membuat tabel

    // Read JSON file
    data, err := os.ReadFile("data/categories.json")  // Membaca file JSON data category
    if err != nil {
        log.Fatal("Error reading categories.json:", err)  // Log error dan hentikan program jika gagal
    }

    // Parse JSON data
    var categories []entity.Category          // Variabel untuk menampung data category dari JSON
    err = json.Unmarshal(data, &categories)   // Mengkonversi JSON ke slice struct Category
    if err != nil {
        log.Fatal("Error parsing categories.json:", err)  // Log error dan hentikan program jika gagal
    }

    // Seed data
    for _, category := range categories {     // Iterasi setiap category dari data JSON
        if err := db.Create(&category).Error; err != nil {  // Menyimpan category ke database
            log.Fatal("Error seeding category:", err)  // Log error dan hentikan program jika gagal
        }
    }

    fmt.Println("🌱 Category data seeded successfully!")  // Pesan sukses seed data
}



// {{{ Penjelasan Fungsi Categories }}}

/*
## Penjelasan Detail
File category.go ini berisi fungsi untuk melakukan seed data category ke dalam database. Berikut penjelasan detailnya:

1. Tujuan : File ini digunakan untuk mengisi database dengan data awal (seed data) untuk entitas Category.
2. Alur Kerja :

    - Menghapus tabel Category yang sudah ada (jika ada)
    - Membuat tabel baru berdasarkan struktur entity Category
    - Membaca data dari file JSON
    - Mengkonversi data JSON ke slice struct Category
    - Menyimpan data Category ke database
3. Fitur Database :

    - DropTable : Menghapus tabel yang sudah ada
    - AutoMigrate : Membuat tabel berdasarkan struktur struct
    - Create : Menyimpan data ke database
4. Penanganan File :

    - Membaca file JSON dari path data/categories.json
    - Menggunakan json.Unmarshal untuk mengkonversi JSON ke struct Go
5. Penanganan Error :

    - Menggunakan log.Fatal untuk menghentikan program jika terjadi error
    - Menampilkan pesan error yang informatif
6. Feedback :

    - Menampilkan pesan sukses dengan emoji untuk setiap tahap proses
    - Memudahkan pengguna untuk memahami progress seed data
Fungsi ini biasanya digunakan saat inisialisasi aplikasi atau untuk keperluan pengujian, di mana database perlu diisi dengan data awal yang konsisten.
*/