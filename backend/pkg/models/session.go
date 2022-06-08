package models

type Session struct {
	ID     string
	UserID string
	// TimeAccessed time.Time
}

type SessionRepository interface {
	// save new session to db
	Set(Session) error
	// read session id from db / compare
	// Get(sID, userID string) error
	// Delete session id from db
	// Delete(sID string) error
}
