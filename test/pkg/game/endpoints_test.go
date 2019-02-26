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

var config *data.Config
var dataStore data.DataStore
var gameService game.Service
var playerService player.Service
var gameController game.Controller

func init() {
	config = data.NewConfiguration()
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

func TestCreateGameReturns201(t *testing.T) {
	expected := models.GameInfo{Name:"Game1", PlayerNames: []string{"Player1", "Player2", "Player3"}}
	rr, err := createGame(expected)
	if err != nil {
		t.Fatal(err)
	}
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("create returned wrong status code: actual %v expected %v", status, http.StatusCreated)
	}
}