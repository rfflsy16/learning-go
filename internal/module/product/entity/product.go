package entity                                // Mendefinisikan package entity untuk modul product

import (
    "time"                                    // Package time untuk tipe data waktu
    
    "github.com/go-playground/validator/v10"  // Package validator untuk validasi data
)

type Product struct {                         // Mendefinisikan struct Product
    ID          uint      `json:"id" gorm:"primaryKey"`  // ID produk sebagai primary key
    Title       string    `json:"title" binding:"max=255"`  // Judul produk dengan validasi maksimal 255 karakter
    Price       float64   `json:"price" binding:"numeric"`  // Harga produk dengan validasi harus numerik
    Description string    `json:"description" binding:"max=255"`  // Deskripsi produk dengan validasi maksimal 255 karakter
    CategoryID  uint      `json:"category_id" gorm:"index"`  // ID kategori sebagai foreign key dengan indeks untuk performa query
    CreatedAt   time.Time `json:"created_at"`  // Waktu pembuatan record
    UpdatedAt   time.Time `json:"updated_at"`  // Waktu pembaruan record
}

func (p *Product) Validate() error {          // Method untuk validasi struct Product
    validate := validator.New()               // Membuat instance validator baru
    return validate.Struct(p)                 // Memvalidasi struct berdasarkan tag binding
}



//  {{{ Penjelasan Struktur Product }}}

/*
package entity                                // Mendefinisikan package entity untuk modul product

import (
    "time"                                    // Package time untuk tipe data waktu
    
    "github.com/go-playground/validator/v10"  // Package validator untuk validasi data
)

type Product struct {                         // Mendefinisikan struct Product
    ID          uint      `json:"id" gorm:"primaryKey"`  // ID produk sebagai primary key
    Title       string    `json:"title" binding:"max=255"`  // Judul produk dengan validasi maksimal 255 karakter
    Price       float64   `json:"price" binding:"numeric"`  // Harga produk dengan validasi harus numerik
    Description string    `json:"description" binding:"max=255"`  // Deskripsi produk dengan validasi maksimal 255 karakter
    CategoryID  uint      `json:"category_id" gorm:"index"`  // ID kategori sebagai foreign key dengan indeks untuk performa query
    CreatedAt   time.Time `json:"created_at"`  // Waktu pembuatan record
    UpdatedAt   time.Time `json:"updated_at"`  // Waktu pembaruan record
}

func (p *Product) Validate() error {          // Method untuk validasi struct Product
    validate := validator.New()               // Membuat instance validator baru
    return validate.Struct(p)                 // Memvalidasi struct berdasarkan tag binding
}
*/