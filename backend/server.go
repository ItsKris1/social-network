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
	mux.HandleFunc("/allUsers", handler.Auth(handler.AllUsers))       // all users + info except current
	mux.HandleFunc("/followers", handler.Auth(handler.GetFollowers))  //follower list
	mux.HandleFunc("/following", handler.Auth(handler.GetFollowing))  // following list
	mux.HandleFunc("/currentUser", handler.Auth(handler.CurrentUser)) //current user data
	mux.HandleFunc("/userData", handler.Auth(handler.UserData))       //userd data based on following status
	mux.HandleFunc("/changeStatus", handler.Auth(handler.UserStatus)) //change status

	mux.HandleFunc("/follow", handler.Auth(handler.Follow)) //follow user
	mux.HandleFunc("/unfollow", handler.Auth(handler.Unfollow))
	mux.HandleFunc("/responseFollowRequest", handler.Auth(handler.ResponseFollowRequest))

	/* ---------------------------------- posts --------------------------------- */
	mux.HandleFunc("/allPosts", handler.Auth(handler.AllPosts))   // all posts- main page
	mux.HandleFunc("/userPosts", handler.Auth(handler.UserPosts)) // all user posts - user page
	mux.HandleFunc("/newPost", handler.Auth(handler.NewPost))     // create route

	/* -------------------------------- comments -------------------------------- */
	mux.HandleFunc("/newComment", handler.Auth(handler.NewComment)) // create route

	/* --------------------------------- groups --------------------------------- */
	mux.HandleFunc("/allGroups", handler.Auth(handler.AllGroups)) // group list

	mux.HandleFunc("/groupInfo", handler.Auth(handler.GroupInfo))         // get group info
	mux.HandleFunc("/groupMembers", handler.Auth(handler.GroupMembers))   // get group members
	mux.HandleFunc("/groupEvents", handler.Auth(handler.GroupEvents))     // get group events
	mux.HandleFunc("/groupPosts", handler.Auth(handler.GroupPosts))       // get group posts
	mux.HandleFunc("/groupRequests", handler.Auth(handler.GroupRequests)) // get group member requests

	mux.HandleFunc("/newGroup", handler.Auth(handler.NewGroup))                           // create new group
	mux.HandleFunc("/newGroupPost", handler.Auth(handler.NewGroupPost))                   // create new group post
	mux.HandleFunc("/newGroupInvite", handler.Auth(handler.NewGroupInvite))               // invite new users to group
	mux.HandleFunc("/newGroupRequest", handler.Auth(handler.NewGroupRequest))             // new request to join a group
	mux.HandleFunc("/responseGroupRequest", handler.Auth(handler.ResponseGroupRequest))   // response to join request
	mux.HandleFunc("/responseInviteRequest", handler.Auth(handler.ResponseInviteRequest)) // response to invite request

	/* --------------------------------- events --------------------------------- */
	// mux.HandleFunc("/newEvent", handler.Auth(handler.NewEvent)) // create new
	// mux.HandleFunc("/participate", handler.Auth(handler.Participate)) // react to participation in event

	return mux
}
