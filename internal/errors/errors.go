package errors

import (
	"encoding/json"
	"net/http"
)

type errBody struct {
	Error string `json:"error"`
}

func write(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(errBody{Error: msg})
}

func BadRequest(w http.ResponseWriter, msg string) {
	write(w, http.StatusBadRequest, msg)
}
func Internal(w http.ResponseWriter, msg string) {
	write(w, http.StatusInternalServerError, msg)
}
