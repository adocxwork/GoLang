package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time

	// Foreign key (belongs to User)
	UserID uint
}
