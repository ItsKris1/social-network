package handlers

import (
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
)

// Find all users and they relation with current user
func (handler *Handler) AllUsers(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	// access user id
	userId := r.Context().Value(utils.UserKey).(string)
	// request all users exccept current + relations
	users, errUsers := handler.repos.UserRepo.GetAllAndFollowing(userId)
	if errUsers != nil {
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	utils.RespondWithUsers(w, users, 200)
}

// Returns user nickname, id and path to avatar
func (handler *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	// access user id
	userId := r.Context().Value(utils.UserKey).(string)
	user, err := handler.repos.UserRepo.GetDataMin(userId)
	if err != nil {
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	utils.RespondWithUsers(w, []models.User{user}, 200)
}
