package handlers

import (
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
)

/* -------------------------------------------------------------------------- */
/*     test handlers / in the futere should hold login / signup and logout    */
/* -------------------------------------------------------------------------- */
/* ----------- not yet connected to front end with right responses ---------- */

// TEST handler for checking if everything connected
func (handler *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello home"))
	// Example for db communication

	// newUser := models.User{Name: "TEST", ID: "5"}
	// err := handler.repos.UserRepo.Add(newUser)
	// if err != nil {
	// 	http.Error(w, "Something went wrong", http.StatusBadRequest)
	// }
}

// TEST handler for protected route
func (handler *Handler) Secret(w http.ResponseWriter, r *http.Request) {
	// access user id
	userId := r.Context().Value(utils.UserKey)
	w.Write([]byte("You have access to secret"))
	fmt.Println(userId, "User validated")
}

// TEST handler for registering user
func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// crete new user instance
	userID := utils.UniqueId()
	newUser := models.User{Name: "User1", ID: userID}

	// Add user in database
	err := handler.repos.UserRepo.Add(newUser)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
	}

	// Start new session for user (Including cookies)
	session := utils.SessionStart(w, r, userID)
	// Save session in database
	errSession := handler.repos.SessionRepo.Set(session)
	if errSession != nil {
		fmt.Println("error on saving cookie in db", errSession)
		return
	}
	w.Write([]byte("Registered successfully"))
}

// TEST  handler for logout
func (handler *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// access user id
	userId := r.Context().Value(utils.UserKey).(string)
	// delete session
	session := models.Session{UserID: userId}
	errSession := handler.repos.SessionRepo.Delete(session)
	if errSession != nil {
		fmt.Println("error on deleting session", errSession)
		return
	}
	// delete cookie
	utils.DeleteCookie(w)
	w.Write([]byte("Logged out"))
}
