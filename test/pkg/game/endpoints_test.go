package game

import (
	"bytes"
	"encoding/json"
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/game"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/auroq/clue-api/pkg/player"
	"github.com/auroq/clue-api/test/pkg/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dataStore data.DataStore
var gameService game.Service
var playerService player.Service
var gameController game.Controller

func init() {
	dataStore = mock.MockMongoDataStore{}
	gameService = game.NewGameService(dataStore)
	playerService = player.NewPlayerService(dataStore)
	gameController = game.NewGameController(gameService, playerService)
}

func createGame(body interface{}) (*httptest.ResponseRecorder, error) {
	jsonValue, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/games", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gameController.CreateGame)
	handler.ServeHTTP(rr, req)
	return rr, nil
}

var createGameStatusTests = []struct {
	name   string
	body   interface{}
	status int
}{
	{"Created", models.GameInfo{Name: "Game1", PlayerNames: []string{"Player1", "Player2", "Player3"}}, 201},
	{"EmptyPlayerNames", models.GameInfo{Name: "Game1", PlayerNames: []string{}}, 400},
	{"NilPlayerNames", models.GameInfo{Name: "Game1", PlayerNames: nil}, 400},
	{"EmptyName", models.GameInfo{Name: "", PlayerNames: []string{"Player1", "Player2", "Player3"}}, 400},
}

func TestCreateGameReturns(t *testing.T) {
	for _, tt := range createGameStatusTests {
		t.Run(tt.name, func(t *testing.T) {
			rr, err := createGame(tt.body)
			if err != nil {
				t.Fatal(err)
			}
			if code := rr.Code; code != tt.status {
				t.Errorf("create returned wrong status code: actual %v expected %v", code, tt.status)
			}
		})
	}
}

func getAllGames() (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", "/games", nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gameController.GetAllGames)
	handler.ServeHTTP(rr, req)
	return rr, nil
}

var getAllGamesStatusTests = []struct {
	name   string
	games  []models.Game
	status int
}{
	{"EmptyList", []models.Game{}, 201},
}

func TestGetAllGamesReturns(t *testing.T) {
	for _, tt := range getAllGamesStatusTests {
		t.Run(tt.name, func(t *testing.T) {
			rr, err := getAllGames()
			if err != nil {
				t.Fatal(err)
			}
			if code := rr.Code; code != tt.status {
				t.Errorf("create returned wrong status code: actual %v expected %v", code, tt.status)
			}
		})
	}
}
