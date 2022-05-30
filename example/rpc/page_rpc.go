package rpc

import (
	"gorm.io/gorm"
	"rhodium"
	"rhodium/example/models"
)

// PostRPC -
type PostRPC struct {
	db *gorm.DB
}

// NewPostRPC -
func NewPostRPC(db *gorm.DB) *PostRPC {
	return &PostRPC{db}
}

// CreatePost -
func (r *PostRPC) CreatePost(ctx rhodium.RPCContext) (map[string]interface{}, error) {
	var post *models.Post
	_ = ctx.Body(&post)

	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"created": post,
	}, nil
}
