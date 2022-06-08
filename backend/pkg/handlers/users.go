package handlers

import (
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
)

// TEST handler for checking if everything connected
func (handler *Handler) Add(w http.ResponseWriter, r *http.Request) {
	// Example for db communication
	newUser := models.User{Name: "TEST", ID: "5"}
	err := handler.repos.UserRepo.Add(newUser)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
	}

}

// TEST handler for checking if everything connected
func (handler *Handler) Test(w http.ResponseWriter, r *http.Request) {
	userID := "5"
	session := utils.SessionStart(w, r, userID)
	err := handler.repos.SessionRepo.Set(session)
	if err != nil {
		fmt.Println("error on saving cookie in db", err)
		return
	}
	fmt.Println("Success! User saved and session created")
}
