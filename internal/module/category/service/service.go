package service                                // Mendefinisikan package service untuk modul category

import (
    "rest-api-go/internal/module/category/entity"  // Mengimpor entity category
    "gorm.io/gorm"                            // Mengimpor ORM GORM
)

type CategoryService struct {                  // Mendefinisikan struct service
    db *gorm.DB                               // Dependency database
}

func NewCategoryService(db *gorm.DB) *CategoryService {  // Constructor untuk service
    return &CategoryService{db}               // Mengembalikan instance service dengan database yang diinjeksi
}

func (s *CategoryService) Create(category *entity.Category) error {  // Method untuk membuat category baru
    if err := category.Validate(); err != nil {  // Validasi data category
        return err                            // Mengembalikan error jika validasi gagal
    }
    return s.db.Create(category).Error        // Menyimpan category ke database dan mengembalikan error jika ada
}

func (s *CategoryService) GetByID(id uint) (*entity.Category, error) {  // Method untuk mendapatkan category berdasarkan ID
    var category entity.Category              // Variabel untuk menampung hasil query
    err := s.db.Preload("Products").First(&category, id).Error  // Query category dengan preload relasi Products
    return &category, err                     // Mengembalikan category dan error jika ada
}

func (s *CategoryService) GetAll() ([]entity.Category, error) {  // Method untuk mendapatkan semua category
    var categories []entity.Category          // Variabel untuk menampung hasil query
    err := s.db.Preload("Products").Find(&categories).Error  // Query semua category dengan preload relasi Products
    return categories, err                    // Mengembalikan categories dan error jika ada
}

func (s *CategoryService) Update(category *entity.Category) error {  // Method untuk memperbarui category
    if err := category.Validate(); err != nil {  // Validasi data category
        return err                            // Mengembalikan error jika validasi gagal
    }

    // Cek apakah category ada
    var existingCategory entity.Category      // Variabel untuk menampung hasil query
    if err := s.db.First(&existingCategory, category.ID).Error; err != nil {  // Query category berdasarkan ID
        return err                            // Mengembalikan error jika category tidak ditemukan
    }

    return s.db.Save(category).Error          // Menyimpan perubahan category ke database dan mengembalikan error jika ada
}

func (s *CategoryService) Delete(id uint) error {  // Method untuk menghapus category
    return s.db.Delete(&entity.Category{}, id).Error  // Menghapus category dari database dan mengembalikan error jika ada
}


// {{{ Penjelasan Fungsi Service }}}

/*
## Penjelasan Detail
File service.go ini berisi implementasi service untuk modul Category. Berikut penjelasan detailnya:

1. Tujuan : File ini berisi logika bisnis untuk operasi CRUD pada entitas Category.
2. Pola Desain :

    - Service Layer : Memisahkan logika bisnis dari handler HTTP
    - Dependency Injection : Database diinjeksi ke dalam service melalui constructor
    - Repository Pattern : Service bertindak sebagai abstraksi untuk akses data
3. Operasi CRUD :

    - Create : Membuat category baru setelah validasi
    - GetByID : Mendapatkan category berdasarkan ID dengan relasi Products
    - GetAll : Mendapatkan semua category dengan relasi Products
    - Update : Memperbarui category setelah validasi dan pengecekan keberadaan
    - Delete : Menghapus category berdasarkan ID
4. Fitur GORM :

    - Preload : Mengambil relasi (Products) bersama dengan data utama
    - First : Mengambil record pertama yang cocok dengan kondisi
    - Find : Mengambil semua record yang cocok dengan kondisi
    - Save : Menyimpan perubahan pada record yang ada
    - Delete : Menghapus record dari database
5. Validasi :

    - Memanggil method Validate() pada entity sebelum operasi Create dan Update
    - Memastikan data valid sebelum berinteraksi dengan database
6. Penanganan Error :

    - Mengembalikan error dari validasi atau operasi database ke handler
    - Memeriksa keberadaan record sebelum update untuk mencegah error
Service ini mengimplementasikan prinsip "fat model, thin controller" di mana logika bisnis berada di service, sementara handler hanya bertanggung jawab untuk menangani HTTP request/response.
*/