package game

import (
	"bytes"
	"encoding/json"
	"github.com/auroq/clue-api/pkg/game"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/auroq/clue-api/pkg/player"
	"github.com/auroq/clue-api/test/pkg/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dataStore mock.MockMongoDataStore
var gameService game.Service
var playerService player.Service
var gameController game.Controller

func init() {
	dataStore = mock.MockMongoDataStore{}
	gameService = game.NewGameService(&dataStore)
	playerService = player.NewPlayerService(&dataStore)
	gameController = game.NewGameController(gameService, playerService)
}

func createGamePost(body interface{}) (*httptest.ResponseRecorder, error) {
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
	{
		"Created",
		models.GameInfo{Name: "Game1", PlayerNames: []string{"Player1", "Player2", "Player3"}},
		http.StatusCreated,
	},
	{
		"EmptyPlayerNames",
		models.GameInfo{Name: "Game1", PlayerNames: []string{}},
		http.StatusBadRequest,
	},
	{
		"NilPlayerNames",
		models.GameInfo{Name: "Game1", PlayerNames: nil},
		http.StatusBadRequest,
	},
	{
		"EmptyName",
		models.GameInfo{Name: "", PlayerNames: []string{"Player1", "Player2", "Player3"}},
		http.StatusBadRequest,
	},
}

func TestCreateGameEndpointReturns(t *testing.T) {
	for _, tt := range createGameStatusTests {
		t.Run(tt.name, func(t *testing.T) {
			rr, err := createGamePost(tt.body)
			if err != nil {
				t.Fatal(err)
			}
			if code := rr.Code; code != tt.status {
				t.Errorf("create returned wrong status code: actual %v expected %v", code, tt.status)
			}
		})
	}
}

func getAllGamesGet() (*httptest.ResponseRecorder, error) {
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
	{"OkEmptyList", []models.Game{}, http.StatusOK},
	{"OkSingleItem", []models.Game{mock.GetGame("Game1", 2)}, http.StatusOK},
	{"OkMultipleGames", []models.Game{mock.GetGame("Game1", 2), mock.GetGame("Game2", 4)}, http.StatusOK},
}

func TestGetAllGamesEndpointReturns(t *testing.T) {
	for _, tt := range getAllGamesStatusTests {
		t.Run(tt.name, func(t *testing.T) {
			var genericGames []interface{}
			for _, game := range tt.games {
				genericGames = append(genericGames, game)
			}
			dataStore.FindValue = genericGames
			rr, err := getAllGamesGet()
			if err != nil {
				t.Fatal(err)
			}
			if code := rr.Code; code != tt.status {
				t.Errorf("create returned wrong status code: actual %v expected %v", code, tt.status)
			}
		})
	}
}
