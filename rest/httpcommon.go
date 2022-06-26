package rest

import (
	"encoding/json"
	"net/http"
)

func SendResponseJSON(w http.ResponseWriter, status int64, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DecodeBodyAsJSON(r *http.Request, structure interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(&structure)
	return err == nil
}
