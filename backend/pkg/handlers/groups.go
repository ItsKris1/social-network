package handlers

import (
	"net/http"
	"social-network/pkg/utils"
)

func (handler *Handler) AllGroups(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	// access user id
	userId := r.Context().Value(utils.UserKey).(string)
	// request all groups + relations
	groups, errGroups := handler.repos.GroupRepo.GetAllAndRelations(userId)
	if errGroups != nil {
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	utils.RespondWithGroups(w, groups, 200)
}
