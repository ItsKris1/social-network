package utils

import (
	"encoding/json"
	"net/http"
	"social-network/pkg/models"
)

type CustomMessage struct {
	Type    string         `json:"type"`              // msg type ex. errorr, success
	Message string         `json:"message,omitempty"` // message itself
	Groups  []models.Group `json:"groups,omitempty"`
	Posts   []models.Post  `json:"posts,omitempty"`
	Users   []models.User  `json:"users,omitempty"`
	// additional properties/ data types can be added
}

// Error takes writer, message, status code and additional error property
// Sets status code in header and encode resp in json
func RespondWithError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	err := CustomMessage{Message: message, Type: "Error"}
	jsonResp, _ := json.Marshal(err)
	w.Write(jsonResp)
}

// responds with success message
func RespondWithSuccess(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	err := CustomMessage{Message: message, Type: "Success"}
	jsonResp, _ := json.Marshal(err)
	w.Write(jsonResp)
}

// responds with success group
func RespondWithGroups(w http.ResponseWriter, groups []models.Group, code int) {
	w.WriteHeader(code)
	err := CustomMessage{Groups: groups, Type: "Success"}
	jsonResp, _ := json.Marshal(err)
	w.Write(jsonResp)
}

// responds with success group
func RespondWithUsers(w http.ResponseWriter, users []models.User, code int) {
	w.WriteHeader(code)
	err := CustomMessage{Users: users, Type: "Success"}
	jsonResp, _ := json.Marshal(err)
	w.Write(jsonResp)
}

// responds with success group
func RespondWithPosts(w http.ResponseWriter, posts []models.Post, code int) {
	w.WriteHeader(code)
	err := CustomMessage{Posts: posts, Type: "Success"}
	jsonResp, _ := json.Marshal(err)
	w.Write(jsonResp)
}
