package main

import (
	"fmt"
	"html/template"
	"net/http"

	"projet-data-DEX/api"
)

func accueil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func combattants(w http.ResponseWriter, r *http.Request) {
	fighters, err := api.GetAllFighters()

	if err != nil {
		http.Error(w, "Erreur API", 502)
		return
	}

	page, err := template.ParseFiles("templates/fighters.html")

	if err != nil {
		http.Error(w, "Erreur HTML", 500)
		return
	}

	err = page.Execute(w, fighters)

	if err != nil {
		fmt.Println(err)
	}
}

func combattant(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Identifiant manquant", 400)
		return
	}

	fighter, err := api.GetFighter(id)

	if err != nil {
		http.Error(w, "Combattant indisponible", 502)
		return
	}

	page, err := template.ParseFiles("templates/fighter.html")

	if err != nil {
		http.Error(w, "Erreur HTML", 500)
		return
	}

	err = page.Execute(w, fighter)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", accueil)
	http.HandleFunc("/fighters", combattants)
	http.HandleFunc("/fighter", combattant)

	fmt.Println("Serveur : http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}

