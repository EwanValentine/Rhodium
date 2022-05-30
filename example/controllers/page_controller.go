package controllers

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"rhodium"
	"rhodium/example/models"
)

// PageController -
type PageController struct {
	db *gorm.DB
}

// NewPageController -
func NewPageController(db *gorm.DB) *PageController {
	return &PageController{db}
}

// Index -
func (c *PageController) Index(ctx rhodium.Context) error {
	return ctx.View("views/index", map[string]interface{}{
		"name": "Ewan",
	})
}

// ListPosts -
func (c *PageController) ListPosts(ctx rhodium.Context) error {
	var posts []*models.Post
	if err := c.db.Find(&posts).Error; err != nil {
		return errors.Wrap(err, "error listing posts")
	}
	return ctx.View("views/posts", map[string]interface{}{
		"posts": posts,
	})
}
