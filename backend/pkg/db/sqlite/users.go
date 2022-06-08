package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type UserRepository struct {
	DB *sql.DB
}

/* --------------- functions that interact with with database --------------- */

// TEST FUNCTION
// Insert new user in db
func (repo *UserRepository) Add(user models.User) error {
	// example code
	stmt, err := repo.DB.Prepare("INSERT INTO users(username, user_id, session_id) values(?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(user.Name, user.ID, "NULL"); err != nil {
		return err
	}
	return nil
}
