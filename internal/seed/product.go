package seed

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "rest-api-go/internal/module/product/entity"

    "gorm.io/gorm"
)

// Products - fungsi untuk seed data product
func Products(db *gorm.DB) {
    // Drop table if exists
    err := db.Migrator().DropTable(&entity.Product{})
    if err != nil {
        log.Fatal("Error dropping table:", err)
    }
    fmt.Println("🗑️  Old product tables dropped successfully")

    // Auto migrate
    err = db.AutoMigrate(&entity.Product{})
    if err != nil {
        log.Fatal("Error migrating product table:", err)
    }
    fmt.Println("📝 New product table created successfully")

    // Read JSON file
    data, err := os.ReadFile("data/products.json")
    if err != nil {
        log.Fatal("Error reading products.json:", err)
    }

    // Parse JSON data
    var products []entity.Product
    err = json.Unmarshal(data, &products)
    if err != nil {
        log.Fatal("Error parsing products.json:", err)
    }

    // Seed data
    for _, product := range products {
        if err := db.Create(&product).Error; err != nil {
            log.Fatal("Error seeding product:", err)
        }
    }

    fmt.Println("🌱 Product data seeded successfully!")
}
