package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type SessionRepository struct {
	DB *sql.DB
}

// insert new session into database
func (repo *SessionRepository) Set(session models.Session) error {
	stmt, errQuery := repo.DB.Prepare("INSERT INTO sessions (session_id, user_id, expiration_time) VALUES (?,?,?)")
	if errQuery != nil {
		return errQuery
	}
	_, err := stmt.Exec(session.ID, session.UserID, session.ExpirationTime)
	if err != nil {
		return err
	}
	return nil
}

// get  session based on session id
func (repo *SessionRepository) Get(sessionID string) (models.Session, error) {
	row := repo.DB.QueryRow("SELECT user_id, expiration_time FROM sessions where session_id = ? LIMIT 1", sessionID)
	var session models.Session
	if err := row.Scan(&session.UserID, &session.ExpirationTime); err != nil {
		return session, err
	}
	session.ID = sessionID
	return session, nil
}

// // delete session from database based on user id
func (repo *SessionRepository) Delete(session models.Session) error {
	stmt, err := repo.DB.Prepare("DELETE FROM sessions WHERE user_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(session.UserID)
	if err != nil {
		return err
	}
	return nil
}

// Update sessions expiration time
func (repo *SessionRepository) Update(session models.Session) error {
	_, err := repo.DB.Exec("UPDATE sessions SET expiration_time = ? WHERE session_id=?", session.ExpirationTime, session.ID)
	if err != nil {
		return err
	}
	return nil
}
