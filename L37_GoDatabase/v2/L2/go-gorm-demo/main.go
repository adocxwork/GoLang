package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-gorm-demo/models"
)

func main() {
	// 1. Connect
	dsn := "host=localhost user=appuser password=pass dbname=appdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	fmt.Println("✅ Connected to Postgres with GORM")

	// 2. Auto-migrate (creates table if not exists)
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("AutoMigrate: %v", err)
	}
	fmt.Println("✅ AutoMigrate complete")

	// 3. Insert user
	user := models.User{Name: "Alice", Email: "alice@example.com"}
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Insert user: %v", err)
	}
	fmt.Printf("Inserted user: %+v\n", user)

	// 4. Query user by email
	var found models.User
	if err := db.First(&found, "email = ?", "alice@example.com").Error; err != nil {
		log.Fatalf("Find user: %v", err)
	}
	fmt.Printf("Found user: %+v\n", found)

	// 5. Update user
	if err := db.Model(&found).Update("Name", "Alice Updated").Error; err != nil {
		log.Fatalf("Update user: %v", err)
	}
	fmt.Printf("Updated user: %+v\n", found)

	// 6. List all users
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("List users: %v", err)
	}
	fmt.Println("All users:")
	for _, u := range users {
		fmt.Printf("  %+v\n", u)
	}

	// 7. Delete a user
	if err := db.Delete(&found).Error; err != nil {
		log.Fatalf("Delete user: %v", err)
	}
	fmt.Println("✅ Deleted user")
}
