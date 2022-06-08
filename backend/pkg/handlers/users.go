package handlers

import (
	"net/http"
)

// TEST handler for checking if everything connected
func (handler *Handler) Add(w http.ResponseWriter, r *http.Request) {
	// Example for db communication
	/*
		newUser := models.User{Name: "TEST"}
		err := handler.repos.UserRepo.Add(newUser)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusBadRequest)
		}
	*/
}
