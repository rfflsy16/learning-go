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
	seed.Products(db)
	seed.Users(db)

	log.Println("✅ All data migrated and seeded successfully")
}
