package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// DB connection string
	dsn := "postgresql://appuser:pass@localhost:5432/appdb?sslmode=disable"

	// Connect to DB with logging enabled
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // logs SQL queries
	})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	fmt.Println("✅ Connected to Postgres")

	// Run migrations (creates tables & relationships)
	err = db.AutoMigrate(&User{}, &Post{}, &Student{}, &Course{})
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	fmt.Println("✅ Migration complete")

	// Insert demo data
	if err := seed(db); err != nil {
		log.Fatalf("seeding failed: %v", err)
	}
	fmt.Println("✅ Seeding complete")

	// --- One-to-Many (User → Posts) ---
	var user User
	db.Preload("Posts").First(&user) // eager load posts
	fmt.Printf("User: %s has posts:\n", user.Name)
	for _, p := range user.Posts {
		fmt.Printf("  - %s\n", p.Title)
	}

	// --- Many-to-Many (Students ↔ Courses) ---
	var student Student
	db.Preload("Courses").First(&student)
	fmt.Printf("Student: %s enrolled in:\n", student.Name)
	for _, c := range student.Courses {
		fmt.Printf("  - %s\n", c.Title)
	}
}
