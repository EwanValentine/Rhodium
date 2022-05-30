package models

import "gorm.io/gorm"

// Post -
type Post struct {
	gorm.Model
	Title string
	Body  string
}
