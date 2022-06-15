package handlers

import (
	"encoding/json"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

func (handler *Handler) Signin(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	if r.Method != "POST" {
		utils.RespondWithError(w, "Error on form submittion", 200)
		return
	}
	/* ---------------------------- read incoming data --------------------------- */
	// Try to decode the JSON request to a LoginUser
	var client models.User
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		utils.RespondWithError(w, "Error on form submittion", 200)
		return
	}
	/* --------------------------- validate user in db -------------------------- */
	// find user with email in db (need Password and user_id)
	dbUser, errDb := handler.repos.UserRepo.FindUserByEmail(client.Email)
	if errDb != nil {
		utils.RespondWithError(w, "Wrong credentials", 200)
		return
	}
	// Compare passwords
	errPwd := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(client.Password))
	if errPwd != nil {
		utils.RespondWithError(w, "Wrong credentials", 200)
		return
	}
	/* ------------------- user valid - create/update session ------------------- */
	// get existing session from db
	_, errSession := handler.repos.SessionRepo.GetByUser(dbUser.ID)
	newSession := utils.SessionStart(w, r, dbUser.ID)
	var errOnSave error
	// Update or create new row in db based on/ if session already exist
	if errSession != nil { // create new session
		errOnSave = handler.repos.SessionRepo.Set(newSession)
	} else { // Update existing
		errOnSave = handler.repos.SessionRepo.Update(newSession)
	}
	if errOnSave != nil {
		utils.RespondWithError(w, "Error on creating new session", 200)
		return
	}
	utils.RespondWithSuccess(w, "Login successful", 200)
}
