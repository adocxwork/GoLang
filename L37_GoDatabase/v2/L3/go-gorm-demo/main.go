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

	// 2. Auto-migrate all models
	if err := db.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Fatalf("AutoMigrate: %v", err)
	}
	fmt.Println("✅ AutoMigrate complete")

	// 3. Seed data (create user + posts)
	user := models.User{
		Name:  "Alice",
		Email: "alice@example.com",
		Posts: []models.Post{
			{Title: "My first post", Content: "Hello world!"},
			{Title: "GORM tips", Content: "How to use GORM with Go"},
		},
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Insert user with posts: %v", err)
	}
	fmt.Printf("Inserted user with posts: %+v\n", user)

	// 4. Query with Preload (load posts automatically)
	var found models.User
	if err := db.Preload("Posts").First(&found, "email = ?", "alice@example.com").Error; err != nil {
		log.Fatalf("Find user with posts: %v", err)
	}

	fmt.Printf("Found user: %s <%s>\n", found.Name, found.Email)
	for _, post := range found.Posts {
		fmt.Printf("  Post: %s - %s\n", post.Title, post.Content)
	}

	// 5. Create a post separately
	newPost := models.Post{
		Title:   "Another post",
		Content: "This was added later",
		UserID:  found.ID,
	}
	if err := db.Create(&newPost).Error; err != nil {
		log.Fatalf("Insert post: %v", err)
	}
	fmt.Println("✅ Added a new post")

	// 6. List all posts
	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		log.Fatalf("List posts: %v", err)
	}
	fmt.Println("All posts:")
	for _, p := range posts {
		fmt.Printf("  %d: %s (UserID=%d)\n", p.ID, p.Title, p.UserID)
	}
}
