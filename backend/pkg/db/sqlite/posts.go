package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type PostRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) GetAll(userID string) ([]models.Post, error) {
	var posts []models.Post
	return posts,nil
}
