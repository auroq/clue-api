package player

import (
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"time"
)

type Service struct {
	client data.DataStore
}

func NewPlayerService(dataStore data.DataStore) Service {
	return Service{dataStore}
}

func (service Service) AddPlayer(name string, human bool) (player models.Player, err error) {
	player = models.Player {
		Name: name,
		Human: human,
		DateCreated: time.Now(),
		DateModified: time.Now(),
	}
	player.ID, err = service.client.Insert("clue-api", "players", player)
	return
}

func (service Service) AddPlayers(names ...string) (players []models.Player, err error) {
	var ps []interface{}
	for _, name := range names {
		player := models.Player{
			ID:           primitive.NewObjectID(),
			Name:         name,
			Human:        true,
			DateCreated:  time.Now(),
			DateModified: time.Now(),
		}
		ps = append(ps, player)
		players = append(players, player)
	}
	_, err = service.client.InsertMany("clue-api", "players", ps...)
	if err != nil {
		return nil, err
	}
	return
}
