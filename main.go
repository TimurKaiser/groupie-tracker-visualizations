package main

import (
	"fmt"
	"log"
	"net/http"
	"server/routes"
)

func main() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./client/style"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./client/js"))))
	routes.InitErrorRoutes()
	routes.Init()
	fmt.Println("Serveur démarré sur http://localhost:3001/")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

// gestion des erreurs
//afficher les dates de concert et emplacement
//toutes l'api doit apparaitre la page web
