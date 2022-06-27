package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type NotifRepository struct {
	DB *sql.DB
}

func (repo *NotifRepository) Save(notification models.Notification) error{
	return nil
}