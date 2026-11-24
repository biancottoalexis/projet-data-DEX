package main

import (
	"fmt"
	"net/http"

	"projet-data-DEX/api"
)

func accueil(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	fighters, err := api.GetAllFighters()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Nombre de combattants :", len(fighters))
	}

	http.HandleFunc("/", accueil)

	fmt.Println("Serveur lancé sur localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
