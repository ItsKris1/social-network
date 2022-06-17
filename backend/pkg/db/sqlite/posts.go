package db

import (
	"database/sql"
	"fmt"
	"social-network/pkg/models"
)

type PostRepository struct {
	DB *sql.DB
}

//Returns all posts for user ->
// group posts if is a member
// all public posts
// Private posts if is a follower
// almost_private if has access
// all posts if user is an author
func (repo *PostRepository) GetAll(userID string) ([]models.Post, error) {
	var posts []models.Post
	rows, err := repo.DB.Query("SELECT post_id , created_by, content, image  FROM posts WHERE (group_id NOT NULL AND (SELECT COUNT() FROM group_users WHERE posts.group_id = group_users.group_id AND '"+userID+"' = group_users.user_id)=1) OR (group_id = NULL  AND visibility = 'PUBLIC') OR (group_id = NULL AND visibility = 'PRIVATE' AND (SELECT COUNT() FROM followers WHERE posts.created_by = followers.user_id AND followers.follower_id = '"+userID+"') = 1 ) OR (group_id = NULL  AND visibility = 'ALMOST_PRIVATE' AND (SELECT COUNT() FROM almost_private WHERE almost_private.post_id = posts.post_id AND almost_private.user_id = '"+userID+"')=1) OR created_by = '"+userID+"' ORDER BY created_at DESC;", userID)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.ID, &post.AuthorID, &post.Content, &post.ImagePath)
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo *PostRepository) New(post models.Post) error {
	stmt, err := repo.DB.Prepare("INSERT INTO posts (post_id, group_id, created_by, content,image,visibility) values (?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("Wrong statement", err)
		return err
	}
	if _, err := stmt.Exec(post.ID, post.GroupID, post.AuthorID, post.Content, post.ImagePath, post.Visibility); err != nil {
		return err
	}
	return nil
}
