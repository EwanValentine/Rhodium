package models

import (
	"gorm.io/gorm"
	"time"
)

// Post -
type Post struct {
	gorm.Model
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
