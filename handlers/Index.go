package handlers

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func getFileNamesWithPrefix(fileNames ...string) []string {
	fileNamesWithPrefix := make([]string, 0, len(fileNames))
	prefix := "./static/"
	for _, fileName := range fileNames {
		fileNamesWithPrefix = append(fileNamesWithPrefix, prefix + fileName)
	}
	return fileNamesWithPrefix
}

func init() {
	prefixedFileNames := getFileNamesWithPrefix("index.html")
	tpl = template.Must(template.ParseFiles(prefixedFileNames...))
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
