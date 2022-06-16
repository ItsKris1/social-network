package handlers

import (
	"net/http"
	"social-network/pkg/utils"
)

func (handler *Handler) AllPosts(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	// access user id
	userId := r.Context().Value(utils.UserKey).(string)
	//request all posts
	posts, errPosts := handler.repos.PostRepo.GetAll(userId)
	if errPosts != nil {
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	for i := 0; i < len(posts); i++ {
		author, err := handler.repos.UserRepo.GetDataMin(posts[i].AuthorID)
		if err != nil {
			utils.RespondWithError(w, "Error on getting data", 200)
			return
		}
		posts[i].Author = author
	}
	utils.RespondWithPosts(w, posts, 200)
}
