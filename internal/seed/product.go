package seed                                  // Mendefinisikan package seed

import (
    "encoding/json"                           // Package untuk encoding/decoding JSON
    "fmt"                                     // Package untuk formatting dan output
    "log"                                     // Package untuk logging
    "os"                                      // Package untuk operasi sistem
    "rest-api-go/internal/module/product/entity"  // Mengimpor entity product

    "gorm.io/gorm"                            // Mengimpor ORM GORM
)

// Products - fungsi untuk seed data product
func Products(db *gorm.DB) {                  // Fungsi untuk seed data product dengan parameter database
    // Drop table if exists
    err := db.Migrator().DropTable(&entity.Product{})  // Menghapus tabel product jika ada
    if err != nil {
        log.Fatal("Error dropping table:", err)  // Log error dan hentikan program jika gagal
    }
    fmt.Println("🗑️  Old product tables dropped successfully")  // Pesan sukses menghapus tabel

    // Auto migrate
    err = db.AutoMigrate(&entity.Product{})  // Membuat tabel product berdasarkan struct entity
    if err != nil {
        log.Fatal("Error migrating product table:", err)  // Log error dan hentikan program jika gagal
    }
    fmt.Println("📝 New product table created successfully")  // Pesan sukses membuat tabel

    // Read JSON file
    data, err := os.ReadFile("data/products.json")  // Membaca file JSON data product
    if err != nil {
        log.Fatal("Error reading products.json:", err)  // Log error dan hentikan program jika gagal
    }

    // Parse JSON data
    var products []entity.Product             // Variabel untuk menampung data product dari JSON
    err = json.Unmarshal(data, &products)     // Mengkonversi JSON ke slice struct Product
    if err != nil {
        log.Fatal("Error parsing products.json:", err)  // Log error dan hentikan program jika gagal
    }

    // Seed data
    for _, product := range products {        // Iterasi setiap product dari data JSON
        if err := db.Create(&product).Error; err != nil {  // Menyimpan product ke database
            log.Fatal("Error seeding product:", err)  // Log error dan hentikan program jika gagal
        }
    }

    fmt.Println("🌱 Product data seeded successfully!")  // Pesan sukses seed data
}



// {{{ Penjelasan Fungsi Products }}}

/*
## Penjelasan Detail
File product.go ini berisi fungsi untuk melakukan seed data product ke dalam database. Berikut penjelasan detailnya:

1. Tujuan : File ini digunakan untuk mengisi database dengan data awal (seed data) untuk entitas Product.
2. Alur Kerja :

    - Menghapus tabel Product yang sudah ada (jika ada)
    - Membuat tabel baru berdasarkan struktur entity Product
    - Membaca data dari file JSON
    - Mengkonversi data JSON ke slice struct Product
    - Menyimpan data Product ke database
3. Fitur Database :

    - DropTable : Menghapus tabel yang sudah ada
    - AutoMigrate : Membuat tabel berdasarkan struktur struct
    - Create : Menyimpan data ke database
4. Penanganan File :

    - Membaca file JSON dari path data/products.json
    - Menggunakan json.Unmarshal untuk mengkonversi JSON ke struct Go
5. Penanganan Error :

    - Menggunakan log.Fatal untuk menghentikan program jika terjadi error
    - Menampilkan pesan error yang informatif
6. Feedback :

    - Menampilkan pesan sukses dengan emoji untuk setiap tahap proses
    - Memudahkan pengguna untuk memahami progress seed data
Fungsi ini sangat mirip dengan fungsi Categories yang telah dijelaskan sebelumnya, menunjukkan konsistensi dalam pendekatan seeding data. Perbedaan utama hanya pada entitas yang digunakan (Product vs Category) dan file JSON yang dibaca.
*/