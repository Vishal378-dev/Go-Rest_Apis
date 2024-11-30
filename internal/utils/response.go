package utils

import (
	"encoding/json"
	"net/http"
)

// data interface{} == data any
func WriteResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
