package data

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"time"
)

type Player struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `json:"name"`
	Human bool `json:"human"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	DateModified time.Time `json:"dateModified,omitempty"`
}

type Game struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `json:"name"`
	Players []primitive.ObjectID `json:"players"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	DateModified time.Time `json:"dateModified,omitempty"`
}
