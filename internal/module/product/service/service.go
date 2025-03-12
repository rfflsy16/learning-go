package service                                // Mendefinisikan package service untuk modul product

import (
    "rest-api-go/internal/module/product/entity"  // Mengimpor entity product
    "gorm.io/gorm"                            // Mengimpor ORM GORM
)

type ProductService struct {                   // Mendefinisikan struct service
    db *gorm.DB                               // Dependency database
}

func NewProductService(db *gorm.DB) *ProductService {  // Constructor untuk service
    return &ProductService{db}                // Mengembalikan instance service dengan database yang diinjeksi
}

func (s *ProductService) Create(product *entity.Product) error {  // Method untuk membuat product baru
    if err := product.Validate(); err != nil {  // Validasi data product
        return err                            // Mengembalikan error jika validasi gagal
    }
    
    // Verify that the category exists
    var count int64                           // Variabel untuk menampung jumlah kategori
    if err := s.db.Model(&entity.Product{}).Where("id = ?", product.CategoryID).Count(&count).Error; err != nil {  // Memeriksa apakah kategori ada
        return err                            // Mengembalikan error jika query gagal
    }
    
    return s.db.Create(product).Error         // Menyimpan product ke database dan mengembalikan error jika ada
}

func (s *ProductService) GetByID(id uint) (*entity.Product, error) {  // Method untuk mendapatkan product berdasarkan ID
    var product entity.Product                // Variabel untuk menampung hasil query
    err := s.db.First(&product, id).Error     // Query product berdasarkan ID
    return &product, err                      // Mengembalikan product dan error jika ada
}

func (s *ProductService) GetAll() ([]entity.Product, error) {  // Method untuk mendapatkan semua product
    var products []entity.Product             // Variabel untuk menampung hasil query
    err := s.db.Find(&products).Error         // Query semua product
    return products, err                      // Mengembalikan products dan error jika ada
}

func (s *ProductService) Update(product *entity.Product) error {  // Method untuk memperbarui product
    if err := product.Validate(); err != nil {  // Validasi data product
        return err                            // Mengembalikan error jika validasi gagal
    }

    // Cek apakah product ada
    var existingProduct entity.Product        // Variabel untuk menampung hasil query
    if err := s.db.First(&existingProduct, product.ID).Error; err != nil {  // Query product berdasarkan ID
        return err                            // Mengembalikan error jika product tidak ditemukan
    }

    return s.db.Save(product).Error           // Menyimpan perubahan product ke database dan mengembalikan error jika ada
}

func (s *ProductService) Delete(id uint) error {  // Method untuk menghapus product
    return s.db.Delete(&entity.Product{}, id).Error  // Menghapus product dari database dan mengembalikan error jika ada
}

func (s *ProductService) GetByCategoryID(categoryID uint) ([]entity.Product, error) {  // Method untuk mendapatkan product berdasarkan CategoryID
    var products []entity.Product             // Variabel untuk menampung hasil query
    err := s.db.Where("category_id = ?", categoryID).Find(&products).Error  // Query product berdasarkan CategoryID
    return products, err                      // Mengembalikan products dan error jika ada
}


// {{{ Penjelasan Fungsi Service }}}

/*
## Penjelasan Detail
File service.go ini berisi implementasi service untuk modul Product. Berikut penjelasan detailnya:

1. Tujuan : File ini berisi logika bisnis untuk operasi CRUD pada entitas Product.
2. Pola Desain :

    - Service Layer : Memisahkan logika bisnis dari handler HTTP
    - Dependency Injection : Database diinjeksi ke dalam service melalui constructor
    - Repository Pattern : Service bertindak sebagai abstraksi untuk akses data
3. Operasi CRUD :

    - Create : Membuat product baru setelah validasi dan verifikasi kategori
    - GetByID : Mendapatkan product berdasarkan ID
    - GetAll : Mendapatkan semua product
    - Update : Memperbarui product setelah validasi dan pengecekan keberadaan
    - Delete : Menghapus product berdasarkan ID
    - GetByCategoryID : Mendapatkan product berdasarkan CategoryID (fitur tambahan)
4. Fitur GORM :

    - First : Mengambil record pertama yang cocok dengan kondisi
    - Find : Mengambil semua record yang cocok dengan kondisi
    - Where : Menentukan kondisi untuk query
    - Save : Menyimpan perubahan pada record yang ada
    - Delete : Menghapus record dari database
5. Validasi :

    - Memanggil method Validate() pada entity sebelum operasi Create dan Update
    - Memastikan data valid sebelum berinteraksi dengan database
6. Penanganan Error :

    - Mengembalikan error dari validasi atau operasi database ke handler
    - Memeriksa keberadaan record sebelum update untuk mencegah error
7. Fitur Tambahan :

    - Method GetByCategoryID memungkinkan pencarian product berdasarkan kategori
    - Ini menunjukkan bagaimana service dapat mengimplementasikan relasi antar entitas
Service ini mengimplementasikan prinsip "fat model, thin controller" di mana logika bisnis berada di service, sementara handler hanya bertanggung jawab untuk menangani HTTP request/response.
*/