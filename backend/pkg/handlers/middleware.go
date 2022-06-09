package handlers

import (
	"context"
	"net/http"
	"social-network/pkg/utils"
	"time"
)

// basic authentification/ check if user logged in
// If not logged in not logged in return
// if logged in continue to handler with user id added to context
// also update expiration time in database
func (handler *Handler) Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get cookie value from request
		sessionId, errCookie := utils.GetCookie(r)
		if errCookie != nil {
			w.Write([]byte("No cookie"))
			return
		}
		// Get session based on session id
		session, errSession := handler.repos.SessionRepo.Get(sessionId)
		if errSession != nil {
			w.Write([]byte("No session"))
			return
		}
		// check if session not expired
		sessionValid := utils.CheckSessionExpiration(session)
		if !sessionValid {
			// if not valid any more delete from db
			handler.repos.SessionRepo.Delete(session)
			// Delete from client browser
			utils.DeleteCookie(w)
			w.Write([]byte("Session expired"))
			return
		} else {
			// Session stil valid -> prolong it by 30 min
			session.ExpirationTime = time.Now().Add(30 * time.Minute)
			handler.repos.SessionRepo.Update(session)
		}
		w.Write([]byte("Auth successful --- "))
		// Auth successful, continue with adding User_id to request context
		ctx := context.WithValue(r.Context(), utils.UserKey, session.UserID)
		next(w, r.WithContext(ctx))
	})
}
