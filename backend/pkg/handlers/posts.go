package handlers

import (
	"net/http"
	"social-network/pkg/models"
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

func (handler *Handler) NewPost(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	if r.Method != "POST" {
		utils.RespondWithError(w, "Error on form submittion", 400)
		return
	}
	err := r.ParseMultipartForm(3145728) // 3MB
	if err != nil {
		utils.RespondWithError(w, "Error in form validation", 400)
		return
	}
	userId := r.Context().Value(utils.UserKey).(string)
	// create new post instance
	newPost := models.Post{
		ID:         utils.UniqueId(),
		Content:    r.PostFormValue("content"),
		GroupID:    r.PostFormValue("group-id"),
		Visibility: r.PostFormValue("visibility"),
		AuthorID:   userId,
	}
	newPost.ImagePath = utils.SaveImage(r)
	// access user id
	errDB := handler.repos.PostRepo.New(newPost)
	if errDB != nil {
		utils.RespondWithError(w, "Error in form validation", 400)
		return
	}
	// save post in db
	utils.RespondWithSuccess(w, "New post created", 200)
}
