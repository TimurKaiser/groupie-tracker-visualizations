package loaders

import (
	"fmt"
	"html/template"
	"net/http"
	"server/constants"
)

type artistData struct {
	Artist    constants.Artist
	Relations constants.Relations
	Dates     constants.Dates
	Locations constants.Locations
}

func Artist(w http.ResponseWriter, r *http.Request, id string, t *template.Template) {
	var fullData artistData

	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_ARTISTS_URL, id), &fullData.Artist)
	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_RELATION_URL, id), &fullData.Relations)
	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_DATES_URL, id), &fullData.Dates)
	GetDataFromURL(fmt.Sprintf("%s/%s", constants.API_LOCATIONS_URL, id), &fullData.Locations)
	pageData := PageData{Data: fullData}
	t.Execute(w, pageData)
}
