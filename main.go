package main

import (
	"encoding/json"
	"fmt"
	"github.com/auroq/clue-api/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/games", GetAllGames).Methods("GET", "OPTIONS")
	router.HandleFunc("/games", CreateGame).Methods("POST", "OPTIONS")

	fmt.Println("Starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var gameInfo struct {
		Name string `json:"name"`
		PlayerNames []string `json:"player_names"`
	}
	err := json.NewDecoder(r.Body).Decode(&gameInfo)
	game, err := data.CreateGame(gameInfo.Name, gameInfo.PlayerNames)
	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode("Invalid request")
		return
	}

	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(game)
}

func GetAllGames(w http.ResponseWriter, r *http.Request) {
	games, err := data.GetAllGames()
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(games)
}
