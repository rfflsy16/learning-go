# Panduan Membuat REST API CRUD dengan Go, MySQL dan GORM

## Struktur Folder dan Penamaan

Go memiliki beberapa pola arsitektur populer. Berikut adalah struktur yang menggunakan pendekatan Clean Architecture/Hexagonal:

```
myapp/
├── cmd/
│   └── api/
│       └── main.go                  # Entry point aplikasi
├── internal/
│   ├── config/                      # Konfigurasi aplikasi
│   │   └── config.go
│   ├── delivery/                    # Layer presentasi (HTTP handlers)
│   │   └── http/
│   │       └── user_handler.go     
│   ├── domain/                      # Entitas dan interface domain
│   │   ├── entity/                  # Entitas bisnis (sebelumnya "model")
│   │   │   └── user.go
│   │   └── repository/              # Interface repository
│   │       └── user_repository.go
│   ├── repository/                  # Implementasi repository 
│   │   └── mysql/
│   │       └── user_repository.go
│   └── usecase/                     # Use cases / business logic (sebelumnya "service")
│       └── user_usecase.go
├── pkg/                             # Package yang bisa digunakan external
│   └── database/
│       └── mysql.go
├── go.mod                           # Dependency management
└── go.sum                           # Dependency lock file
```

## Langkah 1: Inisialisasi Project

```bash
mkdir -p myapp
cd myapp
go mod init github.com/yourusername/myapp
```

## Langkah 2: Instalasi Dependencies

```bash
# Framework web Gin
go get -u github.com/gin-gonic/gin

# GORM dan Driver MySQL
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

# Godotenv untuk environment variables
go get -u github.com/joho/godotenv
```

## Langkah 3: Setup Database Connection

Buat file `pkg/database/mysql.go`:

```go
package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnection membuat dan mengembalikan koneksi database
// Fungsi ini digunakan untuk menginisialisasi koneksi MySQL
func DBConnection() *gorm.DB {
	// Mendapatkan kredensial dari environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Membuat connection string untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		username, password, host, port, dbname)
	
	// Membuka koneksi database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	return db
}
```

## Langkah 4: Buat Entity (Domain Model)

Buat file `internal/domain/entity/user.go`:

```go
package entity

import (
	"gorm.io/gorm"
)

// User merepresentasikan entitas user dalam domain aplikasi
// Struct ini juga digunakan sebagai model GORM untuk pemetaan ke tabel database
type User struct {
	gorm.Model           // Menyertakan ID, CreatedAt, UpdatedAt, DeletedAt
	Name         string  `json:"name" gorm:"type:varchar(100);not null"`       // Nama user
	Email        string  `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"` // Email user (unik)
	Age          int     `json:"age" gorm:"type:int"`                          // Umur user
}

// Catatan:
// - gorm.Model memberikan field ID, CreatedAt, UpdatedAt, DeletedAt secara otomatis
// - Tag `json:"name"` digunakan untuk serialisasi/deserialisasi JSON
// - Tag `gorm:"..."` digunakan untuk konfigurasi kolom database
```

## Langkah 5: Definisikan Repository Interface

Buat file `internal/domain/repository/user_repository.go`:

```go
package repository

import (
	"github.com/yourusername/myapp/internal/domain/entity"
)

// UserRepository mendefinisikan kontrak untuk operasi database terkait user
// Interface ini digunakan untuk memisahkan domain dari implementasi database spesifik
type UserRepository interface {
	FindAll() ([]entity.User, error)     // Mengambil semua user
	FindByID(id uint) (entity.User, error) // Mengambil user berdasarkan ID
	Create(user entity.User) (entity.User, error) // Membuat user baru
	Update(user entity.User) (entity.User, error) // Mengupdate user
	Delete(user entity.User) error      // Menghapus user
}

// Catatan:
// - Interface ini hanya mendefinisikan kontrak, bukan implementasi
// - Ini memungkinkan kita untuk membuat beberapa implementasi (MySQL, PostgreSQL, mock untuk testing)
// - Mengikuti Dependency Inversion Principle dari SOLID
```

## Langkah 6: Implementasi Repository

Buat file `internal/repository/mysql/user_repository.go`:

```go
package mysql

import (
	"github.com/yourusername/myapp/internal/domain/entity"
	"github.com/yourusername/myapp/internal/domain/repository"
	"gorm.io/gorm"
)

// userRepository mengimplementasikan interface UserRepository
// Struct ini bertanggung jawab untuk operasi database terkait user
type userRepository struct {
	db *gorm.DB  // Koneksi database
}

// NewUserRepository membuat instance baru userRepository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

// FindAll mengambil semua user dari database
func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error  // Menggunakan GORM untuk query
	return users, err
}

// FindByID mengambil user berdasarkan ID
func (r *userRepository) FindByID(id uint) (entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error  // First mencari record dengan primary key
	return user, err
}

// Create membuat user baru dalam database
func (r *userRepository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error  // Create menyimpan record baru
	return user, err
}

// Update mengupdate data user yang sudah ada
func (r *userRepository) Update(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error  // Save menyimpan perubahan ke record yang ada
	return user, err
}

// Delete menghapus user dari database
func (r *userRepository) Delete(user entity.User) error {
	return r.db.Delete(&user).Error  // Delete menghapus record
}

// Catatan:
// - Implementasi repository ini menggunakan GORM untuk menyederhanakan operasi database
// - Setiap method mengimplementasikan satu operasi CRUD
```

## Langkah 7: Buat Use Case

Buat file `internal/usecase/user_usecase.go`:

```go
package usecase

import (
	"github.com/yourusername/myapp/internal/domain/entity"
	"github.com/yourusername/myapp/internal/domain/repository"
)

// UserUseCase mendefinisikan kontrak untuk logika bisnis terkait user
type UserUseCase interface {
	GetAllUsers() ([]entity.User, error)        // Mendapatkan semua user
	GetUserByID(id uint) (entity.User, error)   // Mendapatkan user berdasarkan ID
	CreateUser(user entity.User) (entity.User, error) // Membuat user baru
	UpdateUser(user entity.User) (entity.User, error) // Mengupdate user
	DeleteUser(id uint) error                   // Menghapus user berdasarkan ID
}

// userUseCase mengimplementasikan interface UserUseCase
type userUseCase struct {
	userRepo repository.UserRepository  // Dependency ke repository
}

// NewUserUseCase membuat instance baru userUseCase
func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo}
}

// GetAllUsers mengambil semua data user
func (uc *userUseCase) GetAllUsers() ([]entity.User, error) {
	return uc.userRepo.FindAll()
}

// GetUserByID mengambil user berdasarkan ID
func (uc *userUseCase) GetUserByID(id uint) (entity.User, error) {
	return uc.userRepo.FindByID(id)
}

// CreateUser membuat user baru
// Dalam kasus nyata, di sini bisa ada validasi atau logika bisnis lainnya
func (uc *userUseCase) CreateUser(user entity.User) (entity.User, error) {
	return uc.userRepo.Create(user)
}

// UpdateUser mengupdate user yang sudah ada
func (uc *userUseCase) UpdateUser(user entity.User) (entity.User, error) {
	return uc.userRepo.Update(user)
}

// DeleteUser menghapus user berdasarkan ID
func (uc *userUseCase) DeleteUser(id uint) error {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	return uc.userRepo.Delete(user)
}

// Catatan:
// - Use case berisi logika bisnis aplikasi
// - Layer ini tidak tahu tentang HTTP atau database, hanya bekerja dengan domain
// - Use case bergantung pada abstraksi repository, bukan implementasi konkrit
```

## Langkah 8: Buat HTTP Handler

Buat file `internal/delivery/http/user_handler.go`:

```go
package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/myapp/internal/domain/entity"
	"github.com/yourusername/myapp/internal/usecase"
)

// UserHandler menangani HTTP requests terkait user
type UserHandler struct {
	userUseCase usecase.UserUseCase  // Dependency ke use case
}

// NewUserHandler membuat instance baru UserHandler
func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

// GetUsers menangani GET /users
// Handler ini mengambil semua user dan mengembalikannya sebagai JSON
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser menangani GET /users/:id
// Handler ini mengambil user berdasarkan ID dan mengembalikannya sebagai JSON
func (h *UserHandler) GetUser(c *gin.Context) {
	// Parse ID dari URL
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Ambil user dari use case
	user, err := h.userUseCase.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser menangani POST /users
// Handler ini membuat user baru dari data request JSON
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	// Bind JSON request ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat user menggunakan use case
	createdUser, err := h.userUseCase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// UpdateUser menangani PUT /users/:id
// Handler ini mengupdate user berdasarkan ID dan data request JSON
func (h *UserHandler) UpdateUser(c *gin.Context) {
	// Parse ID dari URL
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user entity.User
	// Bind JSON request ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set ID dan update user
	user.ID = uint(id)
	updatedUser, err := h.userUseCase.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser menangani DELETE /users/:id
// Handler ini menghapus user berdasarkan ID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Parse ID dari URL
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Hapus user dengan use case
	if err := h.userUseCase.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// Catatan:
// - Handler bertanggung jawab untuk menangani HTTP requests dan responses
// - Menggunakan Gin framework untuk routing dan handling HTTP
// - Handler tidak berisi logika bisnis, hanya transformasi HTTP ke domain dan sebaliknya
// - Handler bergantung pada use case untuk operasi bisnis
```

## Langkah 9: Buat Config

Buat file `internal/config/config.go`:

```go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig memuat konfigurasi dari file .env
// Fungsi ini membaca variabel lingkungan dari file .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
}

// GetEnv mengambil variabel lingkungan atau mengembalikan nilai default
// Fungsi ini membantu menangani kasus dimana variabel lingkungan tidak ditemukan
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Catatan:
// - godotenv memungkinkan penggunaan file .env untuk konfigurasi
// - Ini memisahkan konfigurasi dari kode (12-Factor App)
```

## Langkah 10: Buat Main Application

Buat file `cmd/api/main.go`:

```go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/myapp/internal/config"
	httpHandler "github.com/yourusername/myapp/internal/delivery/http"
	"github.com/yourusername/myapp/internal/domain/entity"
	mysqlRepo "github.com/yourusername/myapp/internal/repository/mysql"
	"github.com/yourusername/myapp/internal/usecase"
	"github.com/yourusername/myapp/pkg/database"
)

func main() {
	// Load konfigurasi dari file .env
	config.LoadConfig()

	// Koneksi ke database
	db := database.DBConnection()

	// Auto migrate database (membuat tabel berdasarkan struct)
	db.AutoMigrate(&entity.User{})

	// Inisialisasi repository
	userRepository := mysqlRepo.NewUserRepository(db)

	// Inisialisasi use case
	userUseCase := usecase.NewUserUseCase(userRepository)

	// Inisialisasi HTTP handler
	userHandler := httpHandler.NewUserHandler(userUseCase)

	// Setup router Gin
	router := gin.Default()

	// Definisikan routes
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetUsers)           // GET /api/users
			users.GET("/:id", userHandler.GetUser)        // GET /api/users/:id
			users.POST("", userHandler.CreateUser)        // POST /api/users
			users.PUT("/:id", userHandler.UpdateUser)     // PUT /api/users/:id
			users.DELETE("/:id", userHandler.DeleteUser)  // DELETE /api/users/:id
		}
	}

	// Jalankan server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}

// Catatan:
// - main.go adalah entry point aplikasi
// - File ini bertanggung jawab untuk "wiring" semua komponen
// - Dependency Injection dilakukan secara manual di sini
// - Gin digunakan sebagai HTTP router dan framework
```

## Langkah 11: Buat file .env

Buat file `.env` di root project:

```
DB_USERNAME=root
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=myapp_db
PORT=8080
```

## Langkah 12: Buat database MySQL

Buat database MySQL dengan nama `myapp_db`:

```sql
CREATE DATABASE myapp_db;
```

## Langkah 13: Jalankan Aplikasi

```bash
# Navigasi ke folder cmd/api
cd cmd/api

# Jalankan aplikasi
go run main.go
```

## API Endpoints

1. **GET /api/users** - Mendapatkan semua users
2. **GET /api/users/:id** - Mendapatkan user berdasarkan ID
3. **POST /api/users** - Membuat user baru
4. **PUT /api/users/:id** - Mengupdate user 
5. **DELETE /api/users/:id** - Menghapus user

## Penjelasan Penamaan dan Terminologi

### Domain vs Model
- **Domain** lebih tepat dalam konteks Domain-Driven Design (DDD)
- **Entity** lebih baik daripada "Model" karena mewakili entitas bisnis yang sebenarnya
- "Model" sering dikaitkan dengan MVC yang memiliki tanggung jawab luas

### UseCase vs Service
- **UseCase** lebih jelas mewakili kasus penggunaan spesifik dalam aplikasi
- Terminology ini berasal dari Clean Architecture dan lebih menggambarkan tujuan layer ini
- "Service" lebih umum dan kurang spesifik tentang tanggung jawabnya

### Delivery vs Handler
- **Delivery** menunjukkan bahwa ini adalah titik dimana data "dikirimkan" ke pengguna
- Layer ini dapat memiliki berbagai jenis (HTTP, gRPC, CLI) sehingga penamaan lebih umum

### Repository vs DAO
- **Repository** mewakili konsep collection-like interface untuk mengakses domain objects
- Ini lebih fokus pada domain daripada "DAO" (Data Access Object) yang lebih fokus pada database