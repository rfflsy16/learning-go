package database

import (
    "fmt"
    "log"
    "rest-api-go/pkg/config"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func Connect() *gorm.DB {
    // Load config
    cfg := config.LoadConfig()

    // Connect to MariaDB
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    return db
}
