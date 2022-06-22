package handlers

import (
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"
	"strings"
)

/* ------------------------ fetch all posts for user ------------------------ */
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
	for i := 0; i < len(posts); i++ {
		comments, err := handler.repos.CommentRepo.Get(posts[i].ID)
		if err != nil {
			utils.RespondWithError(w, "Error on getting data", 200)
			return
		}
		posts[i].Comments = comments
	}
	utils.RespondWithPosts(w, posts, 200)
}

func (handler *Handler) UserPosts(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	// access current user id
	currentUserId := r.Context().Value(utils.UserKey).(string)

	// get user id from request
	query := r.URL.Query()
	userId := query.Get("id")
	if userId == "" {
		utils.RespondWithError(w, "Error user id", 200)
		return
	}
	//request user posts
	posts, errPosts := handler.repos.PostRepo.GetUserPosts(userId, currentUserId)
	if errPosts != nil {
		fmt.Println("Err: ", errPosts)
		utils.RespondWithError(w, "Error on getting data", 200)
		return
	}
	// Get post author info attached
	for i := 0; i < len(posts); i++ {
		author, err := handler.repos.UserRepo.GetDataMin(posts[i].AuthorID)
		if err != nil {
			utils.RespondWithError(w, "Error on getting data", 200)
			return
		}
		posts[i].Author = author
	}
	// Get comment info for each post
	for i := 0; i < len(posts); i++ {
		comments, err := handler.repos.CommentRepo.Get(posts[i].ID)
		if err != nil {
			utils.RespondWithError(w, "Error on getting data", 200)
			return
		}
		posts[i].Comments = comments
	}
	utils.RespondWithPosts(w, posts, 200)
}

/* ----------------------------- create new post ---------------------------- */
func (handler *Handler) NewPost(w http.ResponseWriter, r *http.Request) {
	w = utils.ConfigHeader(w)
	if r.Method != "POST" {
		utils.RespondWithError(w, "Error on form submittion", 200)
		return
	}
	err := r.ParseMultipartForm(3145728) // 3MB
	if err != nil {
		utils.RespondWithError(w, "Error in form validation", 200)
		return
	}
	// configure data
	userId := r.Context().Value(utils.UserKey).(string)
	visibility := strings.ToUpper(r.PostFormValue("privacy"))
	visibility = strings.Replace(visibility, "-", "_", -1)
	// create new post instance
	newPost := models.Post{
		ID:         utils.UniqueId(),
		Content:    r.PostFormValue("body"),
		GroupID:    r.PostFormValue("group-id"),
		Visibility: visibility,
		AuthorID:   userId,
	}
	// save image in filesystem
	newPost.ImagePath = utils.SaveImage(r)
	// save post in database
	errDB := handler.repos.PostRepo.New(newPost)
	if errDB != nil {
		utils.RespondWithError(w, "Error in form validation", 200)
		return
	}
	// in case of "almost private post", save users with access
	if newPost.Visibility == "ALMOST_PRIVATE" {
		accessList := r.PostFormValue("checkedfollowers")
		fmt.Println("Save: ", accessList)
		for i := 0; i < len(accessList); i++ {
			// save each follower in db
			// fmt.Println("Save: ", accessList[i])
		}

	}
	utils.RespondWithSuccess(w, "New post created", 200)
}
