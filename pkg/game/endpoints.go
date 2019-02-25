package game

import (
	"encoding/json"
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
)

func Endpoints(router *mux.Router) {
	router.HandleFunc("/games", GetAllGames).Methods("GET", "OPTIONS")
	router.HandleFunc("/games", CreateGame).Methods("POST", "OPTIONS")
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var gameInfo models.GameInfo
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
