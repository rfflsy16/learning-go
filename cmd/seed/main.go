package main

import (
	"log"
	"rest-api-go/internal/seed"
	"rest-api-go/pkg/database"
)

func main() {
	// Connect to database
	db := database.Connect()

	// Seed data (termasuk migrasi)
	seed.Users(db)
	seed.Products(db)
	seed.Categories(db)

	log.Println("✅ All data migrated and seeded successfully")
}
