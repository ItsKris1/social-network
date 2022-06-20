package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type CommentRepository struct {
	DB *sql.DB
}

func (repo *CommentRepository) Get(postID string) ([]models.Comment, error) {
	var comments []models.Comment
	rows, err := repo.DB.Query("SELECT comment_id, created_by, (SELECT IFNULL(nickname, first_name || ' ' || last_name) FROM users WHERE users.user_id = comments.created_by), content, image FROM comments WHERE post_id = ? ORDER BY created_at DESC;", postID)
	if err != nil {
		return comments, err
	}
	for rows.Next() {
		var comment models.Comment
		rows.Scan(&comment.ID, &comment.AuthorID, &comment.AuthorNickname, &comment.Content, &comment.ImagePath)
		comments = append(comments, comment)
	}
	return comments, nil
}
