package models

type Comment struct {
	ID string `json:"id"`

	Content        string `json:"content,omitempty"`
	ImagePath      string `json:"image,omitempty"`
	AuthorID       string `json:"author-id,omitempty"`
	AuthorNickname string `json:"author-nickname,omitempty"`
}

type CommentRepository interface {
	// get comment based on postID
	Get(postID string) ([]Comment, error)
}
