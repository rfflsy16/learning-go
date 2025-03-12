package seed

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "rest-api-go/internal/module/category/entity"

    "gorm.io/gorm"
)

// Categories - fungsi untuk seed data category
func Categories(db *gorm.DB) {
    // Drop table if exists
    err := db.Migrator().DropTable(&entity.Category{})
    if err != nil {
        log.Fatal("Error dropping table:", err)
    }
    fmt.Println("🗑️  Old category tables dropped successfully")

    // Auto migrate
    err = db.AutoMigrate(&entity.Category{})
    if err != nil {
        log.Fatal("Error migrating category table:", err)
    }
    fmt.Println("📝 New category table created successfully")

    // Read JSON file
    data, err := os.ReadFile("data/categories.json")
    if err != nil {
        log.Fatal("Error reading categories.json:", err)
    }

    // Parse JSON data
    var categories []entity.Category
    err = json.Unmarshal(data, &categories)
    if err != nil {
        log.Fatal("Error parsing categories.json:", err)
    }

    // Seed data
    for _, category := range categories {
        if err := db.Create(&category).Error; err != nil {
            log.Fatal("Error seeding category:", err)
        }
    }

    fmt.Println("🌱 Category data seeded successfully!")
}
