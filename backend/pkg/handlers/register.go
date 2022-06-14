package handlers

import (
	"fmt"
	"net/http"
	"social-network/pkg/models"
	"social-network/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

/* -------------------------------------------------------------------------- */
/*     test handlers / in the futere should hold login / signup and logout    */
/* -------------------------------------------------------------------------- */
/* ----------- not yet connected to front end with right responses ---------- */

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

// Register user endpont -> validate inputs / save in db / start session
func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
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

	// Create new user instance
	newUser := models.User{
		Email:       r.PostFormValue("email"),
		FirstName:   r.PostFormValue("firstname"),
		LastName:    r.PostFormValue("lastname"),
		Password:    r.PostFormValue("password"),
		Nickname:    r.PostFormValue("nickname"),
		About:       r.PostFormValue("aboutme"),
		DateOfBirth: r.PostFormValue("dateofbirth"),
	}
	// Validate all user fields
	errValid := utils.ValidateNewUser(newUser)
	if errValid != nil {
		utils.RespondWithError(w, "Error in validation", 400)
		return
	}

	// Check if email alredy taken
	if emailUnique, _ := handler.repos.UserRepo.EmailNotTaken(newUser.Email); !emailUnique {
		utils.RespondWithError(w, "Email already taken", 409)
		return
	}
	// Hash password
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPwd)
	// create user id
	userID := utils.UniqueId()
	newUser.ID = userID
	// check if avater added / save in filesystem
	newUser.ImagePath = utils.SaveImage(r)
	// Save user in db
	errSave := handler.repos.UserRepo.Add(newUser)
	if errSave != nil {
		utils.RespondWithError(w, "Couldn't save new user", 500)
		return
	}
	// Start new session for user (Including cookies)
	session := utils.SessionStart(w, r, userID)
	// Save session in database
	errSession := handler.repos.SessionRepo.Set(session)
	if errSession != nil {
		utils.RespondWithError(w, "Error on creating new session", 500)
		return
	}
	/* ------------------------- // Handle image upload ------------------------- */
	// image, h, err := r.FormFile("avatar")
	// if err != nil {
	// 	fmt.Println("Error in geting avatar", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// fmt.Println("avatar: ", h.Filename, image)
}

// TEST  handler for logout
func (handler *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// access user id
	userId := r.Context().Value(utils.UserKey).(string)
	// delete session
	session := models.Session{UserID: userId}
	errSession := handler.repos.SessionRepo.Delete(session)
	if errSession != nil {
		fmt.Println("error on deleting session", errSession)
		return
	}
	// delete cookie
	utils.DeleteCookie(w)
	w.Write([]byte("Logged out"))
}
