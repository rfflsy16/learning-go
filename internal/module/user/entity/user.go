package entity                                // Mendefinisikan package entity untuk modul user

import (
    "time"                                    // Package time untuk tipe data waktu
    
    "github.com/go-playground/validator/v10"  // Package validator untuk validasi data
)

type User struct {                            // Mendefinisikan struct User
    ID          uint      `json:"id" gorm:"primaryKey"`  // ID user sebagai primary key
    Username    string    `json:"username" binding:"max=255"`  // Username user dengan validasi maksimal 255 karakter
    Email       string    `json:"email" binding:"max=255"`  // Email user dengan validasi maksimal 255 karakter
    Password    string    `json:"password" binding:"max=255"`  // Password user dengan validasi maksimal 255 karakter
    CreatedAt   time.Time `json:"created_at"`  // Waktu pembuatan record
    UpdatedAt   time.Time `json:"updated_at"`  // Waktu pembaruan record
}

func (p *User) Validate() error {             // Method untuk validasi struct User
    validate := validator.New()               // Membuat instance validator baru
    return validate.Struct(p)                 // Memvalidasi struct berdasarkan tag binding
}



//  {{{ Penjelasan Struktur User }}}

/*
## Penjelasan Detail
File user.go ini mendefinisikan struktur data untuk entitas User. Berikut penjelasan detailnya:

1. Tujuan : File ini mendefinisikan model data User yang digunakan untuk:

    - Memetakan tabel database ke struct Go (ORM)
    - Mendefinisikan format JSON untuk API
    - Menyediakan validasi data
2. Struktur Data :

    - ID : Primary key untuk user
    - Username : Username user dengan batasan panjang 255 karakter
    - Email : Email user dengan batasan panjang 255 karakter
    - Password : Password user dengan batasan panjang 255 karakter
    - CreatedAt/UpdatedAt : Timestamp untuk audit trail
3. Tag Struct :

    - json : Menentukan nama field dalam respons JSON
    - gorm : Menentukan konfigurasi ORM (primary key)
    - binding : Menentukan aturan validasi
4. Validasi :

    - Method Validate() menggunakan package validator untuk memastikan data valid sebelum disimpan ke database
    - Validasi berdasarkan tag binding pada struct
Entitas User ini merupakan bagian dari pola Repository yang digunakan dalam aplikasi, di mana struct Go digunakan untuk mewakili data dari database dan untuk berinteraksi dengan API.

Dalam pengembangan lebih lanjut, Anda mungkin ingin menambahkan validasi yang lebih ketat untuk email (format email) dan password (kekuatan password), serta menambahkan mekanisme untuk mengenkripsi password sebelum disimpan ke database.
*/