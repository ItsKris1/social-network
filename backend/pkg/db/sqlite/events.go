package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type EventRepository struct {
	DB *sql.DB
}

func (repo *EventRepository) GetAll(groupID string) ([]models.Event, error){
	var events = []models.Event{}
	return events, nil
}