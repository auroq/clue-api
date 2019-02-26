package player

import (
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/models"
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
	return player, err
}
