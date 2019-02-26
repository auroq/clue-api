package game

import (
	"github.com/auroq/clue-api/pkg/game"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/auroq/clue-api/pkg/player"
	"github.com/auroq/clue-api/test/pkg/mock"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"testing"
)

func init() {
	dataStore = mock.MockMongoDataStore{}
	gameService = game.NewGameService(&dataStore)
	playerService = player.NewPlayerService(&dataStore)
	gameController = game.NewGameController(gameService, playerService)
}

var createGameTests = []struct {
	name    string
	players []models.Player
}{
	{
		"OnePlayer",
		[]models.Player{mock.GetPlayer("Player1")},
	},
	{
		"TwoPlayers",
		[]models.Player{mock.GetPlayer("Player2"), mock.GetPlayer("Player3")},
	},
	{
		"ThreePlayers",
		[]models.Player{mock.GetPlayer("Player1"), mock.GetPlayer("Player2"), mock.GetPlayer("Player3")},
	},
}

func TestCreateGame(t *testing.T) {
	for _, tt := range createGameTests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := gameService.CreateGame(tt.name, tt.players)
			if err != nil {
				t.Fatal(err)
			}
			if game.Name != tt.name {
				t.Errorf("create returned wrong name: actual %v expected %v", game.Name, tt.name)
			}
			var blankObjectID primitive.ObjectID
			for i, playerId := range game.Players {
				if playerId == blankObjectID {
					t.Fatal("object id was empty")
				}
				for j, altPlayerId := range game.Players {
					if i == j {
						continue
					} else if playerId == altPlayerId {
						t.Fatal("duplicate object ids found")
					}
				}
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

func TestGetAllGamesReturns(t *testing.T) {
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
