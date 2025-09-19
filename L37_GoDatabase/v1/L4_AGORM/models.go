package main

import "time"

// User model → one user has many posts
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post model → belongs to a user
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string
	UserID    uint      // foreign key (user_id in posts table)
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Student ↔ Course → many-to-many
type Student struct {
	ID      uint     `gorm:"primaryKey"`
	Name    string   `gorm:"not null"`
	Courses []Course `gorm:"many2many:student_courses;"` // join table
}

type Course struct {
	ID       uint      `gorm:"primaryKey"`
	Title    string    `gorm:"not null"`
	Students []Student `gorm:"many2many:student_courses;"`
}
