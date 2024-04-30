package loaders

import (
	"html/template"
	"net/http"
	"server/constants"
)

func Index(w http.ResponseWriter, r *http.Request, t *template.Template) {
	var artists []constants.Artist
	GetDataFromURL(constants.API_ARTISTS_URL, &artists)
	pageData := PageData{Data: artists}
	t.Execute(w, pageData)
}
