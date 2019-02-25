package data

import (
	"context"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"time"
)

var client dataStore

func getClient() dataStore {
	if client == nil {
		client, _ = getConnection()
	}

	return client
}

func addPlayer(name string, human bool) (player models.Player, err error) {
	player = models.Player {
		Name: name,
		Human: human,
		DateCreated: time.Now(),
		DateModified: time.Now(),
	}
	player.ID, err = getClient().insert("clue-api", "players", player)
	return player, err
}

func CreateGame(name string, playerNames []string) (game models.Game, err error) {
	game.Name = name
	for _, playerName := range playerNames {
		player, err := addPlayer(playerName, true)
		if err != nil {
			return game, err
		}
		game.Players = append(game.Players, player.ID)
	}

	game.ID, err = getClient().insert("clue-api", "games", game)
	return game, err
}

func GetAllGames() (games []*models.Game, err error) {
	cur, err := getClient().find("clue-api", "games", bson.D{})
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