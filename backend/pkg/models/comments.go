package models

type Comment struct {
	ID string `json:"id"`

	PostID         string `json:"postId,omitempty"`
	Content        string `json:"content,omitempty"`
	ImagePath      string `json:"image,omitempty"`
	AuthorID       string `json:"authorId,omitempty"`
	AuthorNickname string `json:"authorNickname,omitempty"`
}

type CommentRepository interface {
	// get comment based on postID
	Get(postID string) ([]Comment, error)
	New(Comment) error
}
