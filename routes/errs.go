package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func Error500Handler(w http.ResponseWriter, r *http.Request) {
	pageTemplate, err := template.ParseFiles("client/error500.html")
	if err != nil {
		fmt.Println("Error500Handler", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	pageTemplate.Execute(w, nil)
}

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	pageTemplate, err := template.ParseFiles("client/error404.html")
	if err != nil {
		fmt.Println("Error404Handler", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	pageTemplate.Execute(w, nil)
}

func Error400Handler(w http.ResponseWriter, r *http.Request) {
	pageTemplate, err := template.ParseFiles("client/error400.html")
	if err != nil {
		fmt.Println("Error400Handler", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	pageTemplate.Execute(w, nil)
}

func InitErrorRoutes() {
	http.HandleFunc("/error404", Error404Handler)
	http.HandleFunc("/error500", Error500Handler)
	http.HandleFunc("/error400", Error400Handler)
}
