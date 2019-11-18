package handlers

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Print(err)
	}
	t.Execute(w, nil)
}
