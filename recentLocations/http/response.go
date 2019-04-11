package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"response_code`
}

/*
HTTP Response handling for errors,
Returns valid JSON with error type and response code
*/
func NewErrorResponse(w http.ResponseWriter, statusCode int, response string) {
	error := ErrorResponse{
		true,
		response,
		statusCode,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&error)
	return
}

func NewResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&response)
}

func NewResponseFromJson(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}