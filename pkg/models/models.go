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
	Name         string             `json:"name"`
	Human        bool               `json:"human"`
	DateCreated  time.Time          `json:"dateCreated,omitempty"`
	DateModified time.Time          `json:"dateModified,omitempty"`
}

type Game struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string               `json:"name"`
	Players      []primitive.ObjectID `json:"players"`
	DateCreated  time.Time            `json:"dateCreated,omitempty"`
	DateModified time.Time            `json:"dateModified,omitempty"`
}
