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
	// Seed categories first, then products to maintain foreign key integrity
	seed.Categories(db)
	seed.Products(db)
	seed.Users(db)

	log.Println("✅ All data migrated and seeded successfully")
}
