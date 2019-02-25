package player

import (
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/models"
	"time"
)

type Service struct {
	*data.MongoDataStore
}

func NewPlayerService(dataStore *data.MongoDataStore) Service {
	return Service{dataStore}
}

func (service Service) AddPlayer(name string, human bool) (player models.Player, err error) {
	player = models.Player {
		Name: name,
		Human: human,
		DateCreated: time.Now(),
		DateModified: time.Now(),
	}
	player.ID, err = service.Insert("clue-api", "players", player)
	return player, err
}
