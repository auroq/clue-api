package mock

import (
	"github.com/auroq/clue-api/pkg/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"time"
)

func GetGame(name string, playerNum int) models.Game {
	var playerIds []primitive.ObjectID
	for i := 0; i < playerNum; i++ {
		playerIds = append(playerIds, primitive.NewObjectID())
	}
	return models.Game{
		primitive.NewObjectID(),
		name,
		playerIds,
		time.Now(),
		time.Now(),
	}
}

func GetNonHumanPlayer(name string) models.Player {
	player := GetHumanPlayer(name)
	player.Human = false;
	return player
}

func GetHumanPlayer(name string) models.Player {
	return models.Player{
		primitive.NewObjectID(),
		name,
		true,
		time.Now(),
		time.Now(),
	}
}
