package utils                                 // Mendefinisikan package utils

type Response struct {                        // Mendefinisikan struct Response untuk format respons API
    Success bool        `json:"success"`      // Field untuk menandakan status sukses/gagal, selalu ditampilkan dalam JSON
    Data    interface{} `json:"data,omitempty"`  // Field untuk data respons, tidak ditampilkan jika kosong
    Error   string      `json:"error,omitempty"` // Field untuk pesan error, tidak ditampilkan jika kosong
}

func SuccessResponse(data interface{}) Response {  // Fungsi untuk membuat respons sukses
    return Response{                          // Mengembalikan struct Response
        Success: true,                        // Set Success ke true
        Data:    data,                        // Set Data dengan nilai yang diberikan
    }
}

func ErrorResponse(err string) Response {     // Fungsi untuk membuat respons error
    return Response{                          // Mengembalikan struct Response
        Success: false,                       // Set Success ke false
        Error:   err,                         // Set Error dengan pesan error yang diberikan
    }
}



// {{{ Penjelasan Struktur Response }}}

/*
## Penjelasan Detail
File response.go ini berisi implementasi format respons standar untuk API. Berikut penjelasan detailnya:

1. Tujuan : File ini menyediakan format respons yang konsisten untuk semua endpoint API dalam aplikasi.
2. Struktur Response :

    - Success : Boolean yang menunjukkan apakah request berhasil atau gagal
    - Data : Interface{} yang dapat menampung data respons dalam berbagai bentuk (object, array, string, dll.)
    - Error : String yang berisi pesan error jika request gagal
3. Tag JSON :

    - json:"success" : Field Success selalu ditampilkan dalam respons JSON
    - json:"data,omitempty" : Field Data hanya ditampilkan jika tidak kosong
    - json:"error,omitempty" : Field Error hanya ditampilkan jika tidak kosong
4. Fungsi Helper :

    - SuccessResponse : Membuat respons sukses dengan data yang diberikan
    - ErrorResponse : Membuat respons error dengan pesan error yang diberikan
5. Penggunaan :

    - Fungsi-fungsi ini digunakan di handler untuk mengembalikan respons yang konsisten
    - Contoh: c.JSON(http.StatusOK, utils.SuccessResponse(data))
Format respons yang konsisten ini memudahkan client (frontend) untuk memproses respons API, karena struktur respons selalu sama terlepas dari endpoint yang dipanggil.
*/