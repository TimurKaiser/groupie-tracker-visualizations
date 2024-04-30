package loaders

import (
	"encoding/json"
	"io"
	"net/http"
)

type PageData struct {
	Data any
}

// Recupération des données d'une URL Mashalled ensuite au bon format
func GetDataFromURL(url string, v interface{}) {
	response, err := http.Get(url)
	if err != nil {
		http.RedirectHandler("/error400", http.StatusInternalServerError)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		http.RedirectHandler("/error500", http.StatusInternalServerError)
	}
	err = json.Unmarshal(body, &v)

	if err != nil {
		http.RedirectHandler("/error500", http.StatusInternalServerError)
	}
}
