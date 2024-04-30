package routes

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"server/loaders"
	"strconv"
	"strings"
)

func artist() {
	http.HandleFunc("/artist/", ArtistHandler)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/artist/" {
        http.Redirect(w, r, "/", http.StatusPermanentRedirect)
        return
    }

    id, err := GetArtistsFromURL(r.URL.Path)
    if err != nil {
        http.RedirectHandler("/error500", http.StatusTemporaryRedirect)
        return
        // panic(err)
    }
    //fmt.Println(id)
    str, err1 := strconv.Atoi(id)
    if err1 != nil {
        // panic(err1)
        fmt.Println("Erreur 2")
        http.Redirect(w, r, "/error500", http.StatusTemporaryRedirect)
        return
    }
    fmt.Println(str)
    if str > 52 || str < 1 {
        fmt.Println("Erreur: id de l'artiste invalide")
        http.Redirect(w, r, "/error400", http.StatusTemporaryRedirect)
        return
    } else {
        pageTemplate, err := template.ParseFiles("client/artist.html")
        if err != nil {
            fmt.Println("Erreur 2")
            http.Redirect(w, r, "/error500", http.StatusTemporaryRedirect)
            return
        }
        loaders.Artist(w, r, id, pageTemplate)
    }

}

func GetArtistsFromURL(url string) (string, error) {
	if !(strings.Contains(url, "/artist/")) {
		return "", errors.New("URL is not an artist URL")
	}

	id, found := strings.CutPrefix(url, "/artist/")
	if !found {
		fmt.Println("Erreur 3")
		return "", errors.New("URL is not an artist URL")
	}
	return id, nil
}
