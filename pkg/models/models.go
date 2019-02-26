package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"time"
)

type GameInfo struct {
	Name        string   `json:"name"`
	PlayerNames []string `json:"player_names"`
}

type Player struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `json:primitive.ObjectID"name"`
	Human        bool               `json:"human"`
	DateCreated  time.Time          `json:"dateCreated,omitempty"`
	DateModified time.Time          `json:"dateModified,omitempty"`
}

func (player Player) Equivalent(other Player) bool {
	return player.Name == other.Name &&
		player.Human == other.Human
}

type Game struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string               `json:"name"`
	Players      []primitive.ObjectID `json:"players"`
	DateCreated  time.Time            `json:"dateCreated,omitempty"`
	DateModified time.Time            `json:"dateModified,omitempty"`
}

func (game Game) Equals(other Game) bool {
	return game.ID == other.ID &&
		game.Name == other.Name &&
		game.DateCreated == other.DateCreated &&
		game.DateModified == other.DateModified &&
		func() bool {
			for i, player := range game.Players {
				if player != other.Players[i] {
					return false
				}
			}
			return true
		}()
}
