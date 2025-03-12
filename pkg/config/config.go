package config                                // Mendefinisikan package config

type Config struct {                          // Mendefinisikan struct Config untuk menyimpan konfigurasi aplikasi
    DBHost     string                         // Host database
    DBPort     string                         // Port database
    DBUser     string                         // Username database
    DBPassword string                         // Password database
    DBName     string                         // Nama database
    ServerPort string                         // Port server aplikasi
}

func LoadConfig() *Config {                   // Fungsi untuk memuat konfigurasi
    return &Config{                           // Mengembalikan pointer ke struct Config dengan nilai default
        DBHost:     "localhost",              // Host database default: localhost
        DBPort:     "3306",                   // Port default MariaDB: 3306
        DBUser:     "root",                   // Username database default: root
        DBPassword: "mariadb",                // Password database default: mariadb
        DBName:     "learning-go-DB",         // Nama database default: learning-go-DB
        ServerPort: "8080",                   // Port server default: 8080
    }
}



// {{{ Penjelasan Struktur Config }}}

/*
## Penjelasan Detail
File config.go ini berisi konfigurasi untuk koneksi database dan server aplikasi. Berikut penjelasan detailnya:

1. Tujuan : File ini menyediakan konfigurasi terpusat untuk aplikasi, terutama untuk koneksi database dan pengaturan server.
2. Struktur Config :

    - DBHost : Alamat host database (localhost untuk pengembangan lokal)
    - DBPort : Port database (3306 adalah port default untuk MariaDB/MySQL)
    - DBUser : Username untuk koneksi database
    - DBPassword : Password untuk koneksi database
    - DBName : Nama database yang digunakan aplikasi
    - ServerPort : Port di mana server aplikasi akan berjalan
3. Fungsi LoadConfig :

    - Mengembalikan instance Config dengan nilai default
    - Nilai-nilai ini digunakan untuk koneksi database dan konfigurasi server
4. Penggunaan :

    - Konfigurasi ini biasanya digunakan di main.go untuk menginisialisasi koneksi database dan server
5. Pengembangan Lebih Lanjut :

    - Dalam aplikasi produksi, konfigurasi ini biasanya diambil dari variabel lingkungan atau file konfigurasi eksternal
    - Ini memungkinkan perubahan konfigurasi tanpa perlu mengubah kode
Pendekatan ini memisahkan konfigurasi dari kode aplikasi, yang membuat aplikasi lebih mudah dikonfigurasi untuk lingkungan yang berbeda (pengembangan, pengujian, produksi).
*/