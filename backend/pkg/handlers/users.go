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

// Find all followers
func (handler *Handler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	// access user id
	userId := r.Context().Value(utils.UserKey).(string)
	// request all  following users
	followers, errUsers := handler.repos.UserRepo.GetFollowers(userId)
	if errUsers != nil {
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	utils.RespondWithUsers(w, followers, 200)
}

// Returns user data based on public / private profile and user_id from request
// waits for GET request with query "userId" ->user client is looking for
//  can be used both on own profile and other users
func (handler *Handler) UserData(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	// access user id
	currentUserId := r.Context().Value(utils.UserKey).(string)
	// get user id from request
	query := r.URL.Query()
	userId := query.Get("userId")
	// get if profile public or private
	status, err := handler.repos.UserRepo.ProfileStatus(userId)
	if err != nil {
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	// check if client looking for own profile
	currentUser := (currentUserId == userId)
	var following bool
	if !currentUser {
		// check if current user following user he is looking for
		following, err = handler.repos.UserRepo.IsFollowing(userId, currentUserId)
		if err != nil {
			utils.RespondWithError(w, "Error on getting data", 200)
			return
		}
	}
	// request user info based on following/ and profile status
	// if public or current user or if following  get large data set
	// if private and not following => get small data set
	var user models.User
	if currentUser || following || status == "PUBLIC" { // get full data set
		user, err = handler.repos.UserRepo.GetProfileMax(userId)
	} else {
		user, err = handler.repos.UserRepo.GetProfileMin(userId)
	}
	if err != nil {
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	// tie together stats to user object
	user.Following = following
	user.CurrentUser = currentUser
	user.Status = status

	utils.RespondWithUsers(w, []models.User{user}, 200)
}
