// models/post.go
package models

import "time"

// Post represents a blog post with a title and content
type Post struct {
	ID        int64
	Title     string
	Content   string
	AuthorID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Slug      string
}
