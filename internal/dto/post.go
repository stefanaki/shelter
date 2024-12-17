package dto

import "time"

type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
	UserID  string   `json:"user_id" validate:"required"`
}

type UpdatePostPayload struct {
	Title   string   `json:"title" validate:"max=100"`
	Content string   `json:"content" validate:"max=1000"`
	Tags    []string `json:"tags"`
	UserID  string   `json:"user_id" validate:""`
}
