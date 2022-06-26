package rest

import "net/http"

func root(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello\n Welcome to root endpoint of the service"))
	if err != nil {
		SendResponseJSON(w, http.StatusInternalServerError, err.Error())
	}
}
