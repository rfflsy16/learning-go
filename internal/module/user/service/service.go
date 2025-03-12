package service                                // Mendefinisikan package service untuk modul user

import (
    "rest-api-go/internal/module/user/entity"  // Mengimpor entity user
                                              
    "gorm.io/gorm"                            // Mengimpor ORM GORM
)

type UserService struct {                      // Mendefinisikan struct service
    db *gorm.DB                               // Dependency database
}

func NewUserService(db *gorm.DB) *UserService {  // Constructor untuk service
    return &UserService{db}                   // Mengembalikan instance service dengan database yang diinjeksi
}

func (s *UserService) Create(user *entity.User) error {  // Method untuk membuat user baru
    if err := user.Validate(); err != nil {   // Validasi data user
        return err                            // Mengembalikan error jika validasi gagal
    }
    return s.db.Create(user).Error            // Menyimpan user ke database dan mengembalikan error jika ada
}

func (s *UserService) GetByID(id uint) (*entity.User, error) {  // Method untuk mendapatkan user berdasarkan ID
    var user entity.User                      // Variabel untuk menampung hasil query
    err := s.db.First(&user, id).Error        // Query user berdasarkan ID
    return &user, err                         // Mengembalikan user dan error jika ada
}

func (s *UserService) GetAll() ([]entity.User, error) {  // Method untuk mendapatkan semua user
    var users []entity.User                   // Variabel untuk menampung hasil query
    err := s.db.Find(&users).Error            // Query semua user
    return users, err                         // Mengembalikan users dan error jika ada
}

func (s *UserService) Update(user *entity.User) error {  // Method untuk memperbarui user
    if err := user.Validate(); err != nil {   // Validasi data user
        return err                            // Mengembalikan error jika validasi gagal
    }

    // Cek apakah user ada
    var existingUser entity.User              // Variabel untuk menampung hasil query
    if err := s.db.First(&existingUser, user.ID).Error; err != nil {  // Query user berdasarkan ID
        return err                            // Mengembalikan error jika user tidak ditemukan
    }

    return s.db.Save(user).Error              // Menyimpan perubahan user ke database dan mengembalikan error jika ada
}

func (s *UserService) Delete(id uint) error {  // Method untuk menghapus user
    return s.db.Delete(&entity.User{}, id).Error  // Menghapus user dari database dan mengembalikan error jika ada
}


// {{{ Penjelasan Fungsi Service }}}

/*
## Penjelasan Detail
File service.go ini berisi implementasi service untuk modul User. Berikut penjelasan detailnya:

1. Tujuan : File ini berisi logika bisnis untuk operasi CRUD pada entitas User.
2. Pola Desain :

    - Service Layer : Memisahkan logika bisnis dari handler HTTP
    - Dependency Injection : Database diinjeksi ke dalam service melalui constructor
    - Repository Pattern : Service bertindak sebagai abstraksi untuk akses data
3. Operasi CRUD :

    - Create : Membuat user baru setelah validasi
    - GetByID : Mendapatkan user berdasarkan ID
    - GetAll : Mendapatkan semua user
    - Update : Memperbarui user setelah validasi dan pengecekan keberadaan
    - Delete : Menghapus user berdasarkan ID
4. Fitur GORM :

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