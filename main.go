package main

import (
	"fmt"
	"net/http"
)

func accueil(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	http.HandleFunc("/", accueil)

	fmt.Println("Serveur lancé sur localhost:8080")
	http.ListenAndServe(":8080", nil)
}
