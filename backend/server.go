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
	mux.HandleFunc("/allUsers", handler.Auth(handler.AllUsers))
	mux.HandleFunc("/getFollowers", handler.Auth(handler.GetFollowers))
	mux.HandleFunc("/currentUser", handler.Auth(handler.CurrentUser))

	/* ---------------------------------- posts --------------------------------- */
	mux.HandleFunc("/allPosts", handler.Auth(handler.AllPosts))
	mux.HandleFunc("/newPost", handler.Auth(handler.NewPost))

	/* --------------------------------- groups --------------------------------- */
	mux.HandleFunc("/allGroups", handler.Auth(handler.AllGroups))
	return mux
}
