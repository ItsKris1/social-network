package models

type Post struct {
	ID         string `json:"id"`
	Content    string `json:"content,omitempty"`
	ImagePath  string `json:"image,omitempty"`
	AuthorID   string `json:"authorId,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	GroupID    string `json:"groupId,omitempty"`
	// for sending back with author
	Author   User      `json:"author,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

type PostRepository interface {
	// Get all posts that user have access to
	GetAll(userID string) ([]Post, error)
	// get user posts that current user have access to
	GetUserPosts(userID, currentUserID string) ([]Post, error)
	New(Post) error
}
