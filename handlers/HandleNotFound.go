package handlers

import "net/http"

// Handler for invalid route
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	WriteNotFoundMessage(w, r.URL.Path)
}

