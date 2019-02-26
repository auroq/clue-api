package game

import (
	"context"
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"time"
)

type Service struct {
	client data.DataStore
}

func NewGameService(dataStore data.DataStore) Service {
	return Service{dataStore}
}

func (service Service) CreateGame(name string, players []models.Player) (game models.Game, err error) {
	game.Name = name
	for _, player := range players {
		game.Players = append(game.Players, player.ID)
	}

	game.DateCreated = time.Now()
	game.DateModified = time.Now()
	game.ID, err = service.client.Insert("clue-api", "games", game)
	return game, err
}

func (service Service) GetAllGames() (games []*models.Game, err error) {
	cur, err := service.client.Find("clue-api", "games", bson.D{})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var game models.Game
		err := cur.Decode(&game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	_ = cur.Close(context.TODO())

	return games, nil
}