package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structure de données pour stocker les informations sur un Pokémon
type Pokemon struct {
	Name string `json:"name"` // Nom du Pokémon
}

func accueil(w http.ResponseWriter, r *http.Request) {
	// Affiche la page d'accueil
	http.ServeFile(w, r, "accueil.html")
}

func recherche(w http.ResponseWriter, r *http.Request) {
	// Affiche la page de recherche
	http.ServeFile(w, r, "recherche.html")
}

func main() {
	// Définit les routes de votre application
	http.HandleFunc("/", accueil)
	http.HandleFunc("/recherche", recherche)
	http.HandleFunc("/pokemon", getPokemon) // Ajouter cette ligne pour la route "/pokemon"

	// Démarre le serveur sur le port 8080
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}

// Fonction de gestion de l'endpoint "/pokemon"
func getPokemon(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID du Pokémon à partir des paramètres de la requête GET
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintf(w, "Veuillez entrer l'ID d'un Pokémon à rechercher.")
		return
	}

	// Envoie une requête HTTP GET à l'API PokeAPI pour récupérer les informations sur le Pokémon correspondant à l'ID
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id))
	if err != nil {
		fmt.Fprintf(w, "Erreur : %v", err)
		return
	}
	defer resp.Body.Close()

	// Décode les données JSON renvoyées par l'API PokeAPI dans une structure de données Pokemon
	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		fmt.Fprintf(w, "Erreur : %v", err)
		return
	}

	// Affiche le nom du Pokémon sur une page HTML
	fmt.Fprintf(w, "<html><body><h1>%s</h1></body></html>", pokemon.Name)
}
