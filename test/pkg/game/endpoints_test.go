package game

import (
	"bytes"
	"encoding/json"
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/game"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/auroq/clue-api/pkg/player"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGameWithValidData(t *testing.T) {
	expected := models.GameInfo{Name:"Game1", PlayerNames: []string{"Player1", "Player2", "Player3"}}
	jsonValue, _ := json.Marshal(expected)
	req, err := http.NewRequest("POST", "/games", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	config := data.NewConfiguration()
	dataStore, _ := data.NewDbConnection(config)
	gameService := game.NewGameService(&dataStore)
	playerService := player.NewPlayerService(&dataStore)
	gameController := game.NewGameController(gameService, playerService)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gameController.CreateGame)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("create returned wrong status code: actual %v expected %v", status, http.StatusCreated)
	}
}