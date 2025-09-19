package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DSN same as before
	dsn := "postgresql://appuser:pass@localhost:5432/appdb?sslmode=disable"

	// Open GORM DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	fmt.Println("✅ Connected to Postgres via GORM")

	// Auto-migrate: creates table if not exists (or updates schema)
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed migration: %v", err)
	}
	fmt.Println("✅ Migration complete")

	// Create
	user := User{Name: "Alice GORM", Email: "alicegorm@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf("failed to insert user: %v", result.Error)
	}
	fmt.Printf("Inserted user: %+v\n", user)

	// Query (find by primary key)
	var fetched User
	db.First(&fetched, user.ID) // SELECT * FROM users WHERE id=... LIMIT 1
	fmt.Printf("Fetched by ID: %+v\n", fetched)

	// Query (find by email)
	var byEmail User
	db.First(&byEmail, "email = ?", "alicegorm@example.com")
	fmt.Printf("Fetched by Email: %+v\n", byEmail)

	// Update
	db.Model(&fetched).Update("Name", "Alice Updated")
	fmt.Println("✅ Updated name")

	// Delete
	db.Delete(&fetched)
	fmt.Println("✅ Deleted user")
}
