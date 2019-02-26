package player

import (
	"github.com/auroq/clue-api/pkg/game"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/auroq/clue-api/pkg/player"
	"github.com/auroq/clue-api/test/pkg/mock"
	"golang.org/x/crypto/openpgp/errors"
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

var addPlayerTests = []struct {
	name   string
	player models.Player
	err    error
}{
	{
		"Human",
		mock.GetHumanPlayer("Player1"),
		nil,
	},
	{
		"HumanEmptyName",
		mock.GetHumanPlayer(""),
		errors.InvalidArgumentError("player name cannot be empty"),
	},
	{
		"NonHuman",
		mock.GetNonHumanPlayer("Player1"),
		nil,
	},
	{
		"NonHumanEmptyName",
		mock.GetNonHumanPlayer(""),
		errors.InvalidArgumentError("player name cannot be empty"),
	},
}

func TestAddPlayer(t *testing.T) {
	for _, tt := range addPlayerTests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.player
			actual, err := playerService.AddPlayer(expected.Name, expected.Human)
			if err != nil {
				if err == tt.err {
					return
				}
				t.Fatal(err)
			}
			if !actual.Equivalent(expected) {
				t.Fatal("actual player is not equivalent to expected")
			}
		})
	}
}

var getAllGamesTests = []struct {
	name  string
	games []models.Game
}{
	{"EmptyList", []models.Game{}},
	{"OkEmptyList", []models.Game{}},
	{"OkSingleItem", []models.Game{mock.GetGame("Game1", 2)}},
	{"OkMultipleGames", []models.Game{mock.GetGame("Game1", 2), mock.GetGame("Game2", 4)}},
}

func TestGetAllGamesWith(t *testing.T) {
	for _, tt := range getAllGamesTests {
		t.Run(tt.name, func(t *testing.T) {
			var genericGames []interface{}
			for _, game := range tt.games {
				genericGames = append(genericGames, game)
			}
			dataStore.FindValue = genericGames
			response, err := gameService.GetAllGames()
			if err != nil {
				t.Fatal(err)
			}
			for i, expected := range tt.games {
				actual := response[i]
				if !expected.Equals(actual) {
					t.Fatal("game list returned did not match expected")
				}
			}
		})
	}
}
