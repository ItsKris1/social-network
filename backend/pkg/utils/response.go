package utils

import (
	"encoding/json"
	"net/http"
)

type CustomMessage struct {
	Type    string `json:"type"`    // msg type ex. errorr, success
	Message string `json:"message"` // message itself
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
