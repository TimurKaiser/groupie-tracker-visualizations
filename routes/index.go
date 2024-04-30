package routes

import (
	"html/template"
	"net/http"
	"server/loaders"
	"strings"
)

func index() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	idPath := "/artist/"
	if r.URL.Path != "/" && !strings.HasPrefix(r.URL.Path, idPath) && r.URL.Path != idPath {
		http.Redirect(w, r, "/error404", http.StatusTemporaryRedirect)
		return
	}
	pageTemplate, err := template.ParseFiles("client/index.html")
	if err != nil {
		http.Redirect(w, r, "/error500", http.StatusTemporaryRedirect)
	}
	loaders.Index(w, r, pageTemplate)
}
