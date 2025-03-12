package entity                                  // Mendefinisikan package entity untuk modul category

import (
    "rest-api-go/internal/module/product/entity"  // Mengimpor entity product untuk relasi
    "time"                                         // Package time untuk tipe data waktu
    
    "github.com/go-playground/validator/v10"      // Package validator untuk validasi data
)

type Category struct {                           // Mendefinisikan struct Category
    ID          uint                `json:"id" gorm:"primaryKey"`  // ID kategori sebagai primary key
    Name        string              `json:"name" binding:"max=255"`  // Nama kategori dengan validasi maksimal 255 karakter
    Products    []entity.Product    `json:"products,omitempty" gorm:"foreignKey:CategoryID"`  // Relasi one-to-many dengan Product
    CreatedAt   time.Time           `json:"created_at"`  // Waktu pembuatan record
    UpdatedAt   time.Time           `json:"updated_at"`  // Waktu pembaruan record
}

func (p *Category) Validate() error {           // Method untuk validasi struct Category
    validate := validator.New()                 // Membuat instance validator baru
    return validate.Struct(p)                   // Memvalidasi struct berdasarkan tag binding
}


//  {{{ Penjelasan Struktur Category }}}

/*
## Penjelasan Detail
File category.go ini mendefinisikan struktur data untuk entitas Category. Berikut penjelasan detailnya:

1. Tujuan : File ini mendefinisikan model data Category yang digunakan untuk:

    - Memetakan tabel database ke struct Go (ORM)
    - Mendefinisikan format JSON untuk API
    - Menyediakan validasi data
2. Struktur Data :

    - ID : Primary key untuk kategori
    - Name : Nama kategori dengan batasan panjang 255 karakter
    - Products : Relasi one-to-many dengan entitas Product
    - CreatedAt/UpdatedAt : Timestamp untuk audit trail
3. Tag Struct :

    - json : Menentukan nama field dalam respons JSON
    - gorm : Menentukan konfigurasi ORM (primary key, foreign key)
    - binding : Menentukan aturan validasi
4. Relasi Database :

    - Category memiliki relasi one-to-many dengan Product
    - gorm:"foreignKey:CategoryID" menentukan bahwa field CategoryID di tabel Product adalah foreign key yang merujuk ke tabel Category
5. Validasi :

    - Method Validate() menggunakan package validator untuk memastikan data valid sebelum disimpan ke database
    - Validasi berdasarkan tag binding pada struct
## Konsep Penting
1. ORM (Object-Relational Mapping) : GORM digunakan untuk memetakan struct Go ke tabel database tanpa perlu menulis SQL secara manual.
2. Struct Tags : Tag seperti json , gorm , dan binding memberikan metadata tambahan pada field struct yang digunakan oleh berbagai library.
3. Relasi Antar Entitas : Relasi one-to-many antara Category dan Product diimplementasikan dengan slice Products dan foreign key.
4. Validasi Data : Validasi dilakukan pada level entity untuk memastikan integritas data sebelum berinteraksi dengan database.
Entitas ini adalah dasar dari pola Repository yang digunakan dalam aplikasi, di mana struct Go digunakan untuk mewakili data dari database dan untuk berinteraksi dengan API.
*/