package game

import (
	"encoding/json"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/auroq/clue-api/pkg/player"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	gameService   Service
	playerService player.Service
}

func NewGameController(gameService Service, playerService player.Service) Controller {
	return Controller{
		gameService:   gameService,
		playerService: playerService,
	}
}

func (controller Controller) Endpoints(router *mux.Router) {
	router.HandleFunc("/games", controller.GetAllGames).Methods("GET", "OPTIONS")
	router.HandleFunc("/games", controller.CreateGame).Methods("POST", "OPTIONS")
}

func respond(w http.ResponseWriter, statusCode int, response interface{}) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}

func (controller Controller) CreateGame(w http.ResponseWriter, r *http.Request) {
	var gameInfo models.GameInfo
	err := json.NewDecoder(r.Body).Decode(&gameInfo)

	if len(gameInfo.PlayerNames) <= 0 || len(gameInfo.Name) <= 0 {
		respond(w, http.StatusBadRequest, "Invalid request")
		return
	}

	players, err := controller.playerService.AddPlayers(gameInfo.PlayerNames...)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	game, err := controller.gameService.CreateGame(gameInfo.Name, players)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	respond(w, http.StatusCreated, game)
}

func (controller Controller) GetAllGames(w http.ResponseWriter, r *http.Request) {
	games, err := controller.gameService.GetAllGames()
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}
	respond(w, http.StatusOK, games)
}
