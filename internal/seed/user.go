package seed

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "rest-api-go/internal/module/user/entity"

    "gorm.io/gorm"
)

// Users - fungsi untuk seed data user
func Users(db *gorm.DB) {
    // Drop table if exists
    err := db.Migrator().DropTable(&entity.User{})
    if err != nil {
        log.Fatal("Error dropping table:", err)
    }
    fmt.Println("🗑️  Old user tables dropped successfully")

    // Auto migrate
    err = db.AutoMigrate(&entity.User{})
    if err != nil {
        log.Fatal("Error migrating user table:", err)
    }
    fmt.Println("📝 New user table created successfully")

    // Read JSON file
    data, err := os.ReadFile("data/users.json")
    if err != nil {
        log.Fatal("Error reading users.json:", err)
    }

    // Parse JSON data
    var users []entity.User
    err = json.Unmarshal(data, &users)
    if err != nil {
        log.Fatal("Error parsing users.json:", err)
    }

    // Seed data
    for _, user := range users {
        if err := db.Create(&user).Error; err != nil {
            log.Fatal("Error seeding user:", err)
        }
    }

    fmt.Println("🌱 User data seeded successfully!")
}
