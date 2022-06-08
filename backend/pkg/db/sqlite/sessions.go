package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type SessionRepository struct {
	DB *sql.DB
}

// insert new session id into database
func (repo *SessionRepository) Set(session models.Session) error {
	_, err := repo.DB.Exec("UPDATE Users SET session_id = ? WHERE user_id=?", session.ID, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

// // get Session based on session id
// func (repo *SessionRepository) Get(sID string) (models.Session, error) {

// }
// // delete session from database based on id
// func (repo *SessionRepository) Delete(sID string) error {

// }
