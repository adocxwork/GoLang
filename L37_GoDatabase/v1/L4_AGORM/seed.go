package main

import (
	"gorm.io/gorm"
)

func seed(db *gorm.DB) error {
	// Create a user with posts
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Posts: []Post{
			{Title: "First Post", Content: "Hello World"},
			{Title: "Second Post", Content: "Learning GORM"},
		},
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	// Create many-to-many students & courses
	students := []Student{
		{Name: "John", Courses: []Course{{Title: "Math"}, {Title: "Physics"}}},
		{Name: "Emma", Courses: []Course{{Title: "Math"}, {Title: "Biology"}}},
	}
	if err := db.Create(&students).Error; err != nil {
		return err
	}

	return nil
}
