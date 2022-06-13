package utils

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	Message  string `json:"message"`
	// Property can be used for targeting specific fields
	Property string `json:"property"`
}
// Error takes writer, message, status code and additional error property
// Sets status code in header and encode resp in json
func RespondWithError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	err := CustomError{Message: message}
	jsonResp, _ := json.Marshal(err)
	w.Write(jsonResp)
}
