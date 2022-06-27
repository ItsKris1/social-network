package main

import (
	"fmt"
	"net/http"
	sqlite "social-network/pkg/db/sqlite"
	"social-network/pkg/handlers"
	"social-network/pkg/utils"
)

func main() {
	// initialize database
	db := sqlite.InitDB()
	defer db.Close()
	// temp
	// initialize repositories
	repos := sqlite.InitRepositories(db)
	// initialize handlers with connection to repositories
	handler := handlers.InitHandlers(repos)

	// set up server address and routes
	server := &http.Server{
		Addr:    ":8081",
		Handler: setRoutes(handler),
	}

	fmt.Printf("Server started at http://localhost" + server.Addr + "\n")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error", err)
	}
}

// Set up all routes
func setRoutes(handler *handlers.Handler) http.Handler {
	mux := http.NewServeMux()
	/* ------------------------------ image server ------------------------------ */
	fs := http.FileServer(http.Dir("./imageUpload"))
	mux.Handle("/imageUpload/", http.StripPrefix("/imageUpload/", utils.ConfigFSHeader(fs)))
	/* ------------------------------- auth route ------------------------------- */
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/signin", handler.Signin)
	mux.HandleFunc("/logout", handler.Auth(handler.Logout))
	mux.HandleFunc("/sessionActive", handler.SessionActive)

	/* ---------------------------------- users --------------------------------- */
	mux.HandleFunc("/allUsers", handler.Auth(handler.AllUsers)) // all users + info except current
	mux.HandleFunc("/followers", handler.Auth(handler.GetFollowers)) //follower list
	mux.HandleFunc("/following", handler.Auth(handler.GetFollowing)) // following list
	mux.HandleFunc("/currentUser", handler.Auth(handler.CurrentUser)) //current user data
	mux.HandleFunc("/userData", handler.Auth(handler.UserData)) //userd data based on following status
	mux.HandleFunc("/changeStatus", handler.Auth(handler.UserStatus)) //change status

	/* ---------------------------------- posts --------------------------------- */
	mux.HandleFunc("/allPosts", handler.Auth(handler.AllPosts)) // all posts- main page
	mux.HandleFunc("/userPosts", handler.Auth(handler.UserPosts)) // all user posts - user page
	mux.HandleFunc("/newPost", handler.Auth(handler.NewPost)) // create route

	/* -------------------------------- comments -------------------------------- */
	mux.HandleFunc("/newComment", handler.Auth(handler.NewComment)) // create route

	/* --------------------------------- groups --------------------------------- */
	mux.HandleFunc("/allGroups", handler.Auth(handler.AllGroups)) // group list
	return mux
}
