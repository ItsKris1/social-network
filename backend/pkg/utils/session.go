package utils

import (
	"net/http"
	"social-network/pkg/models"

	uuid "github.com/satori/go.uuid"
)

/* ---- session and cookie funcionality communicating with client request --- */
/* -------------------------------------------------------------------------- */
/*                                   session                                  */
/* -------------------------------------------------------------------------- */

// Creates creates sid, http cookie, sends it to client and returns session
func SessionStart(w http.ResponseWriter, r *http.Request, userID string) models.Session {
	sessionID := CreateSessionId()
	// create cookie
	cookie := CreateCookie(sessionID)
	// create session
	session := models.Session{
		ID:     sessionID,
		UserID: userID,
		// TimeAccessed: time.Now(),
	}
	// Send cookie to client
	http.SetCookie(w, &cookie)
	// return session
	return session
}

/* -------------------------------------------------------------------------- */
/*                                   cookie                                   */
/* -------------------------------------------------------------------------- */

// session cookie blueprint
func CreateCookie(sessionID string)http.Cookie{
	return http.Cookie{
		Name:     "session-id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   3600, //1h
	}
}

func CreateSessionId() string {
	return uuid.NewV4().String()
}
