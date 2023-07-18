package comment

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Defines our comment structure
type Comment struct {
	gorm.Model 		// Tell GORM this will be structured in the DB
	Slug string		// Represents the path where this comment is posted
	Body string		// Conetent of the comment
	Author string
	CreatedAt time.Time
}

// CommentService - the interface for our comment service
type CommentService interface{
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

