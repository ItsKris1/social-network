package handlers

import (
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
)

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
		fmt.Println("error on saving cookie in db", err)
		return
	}
	w.Write([]byte("Registered successfully"))
	fmt.Println("Success! User saved and session created")
}
