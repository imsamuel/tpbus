package handlers

import "net/http"

// Handler for resource that cannot be found
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	notFoundMessage, err := getNotFoundMessage(r.URL.Path)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(notFoundMessage)
}

