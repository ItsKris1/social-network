package models

type Post struct {
	ID        string `json:"id"`
	Content   string `json:"content,omitempty"`
	ImagePath string `json:"image,omitempty"`
	AuthorID  string
	Author    User `json:"author,omitempty"`
}

type PostRepository interface {
	// Get all posts that user have access to
	GetAll(userID string) ([]Post, error)
}
