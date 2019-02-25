package main

import (
	"time"
)

type Player struct {
	ID interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `json:"name"`
	Human bool `json:"human"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	DateModified time.Time `json:"dateModified,omitempty"`
}

type Game struct {
	ID interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `json:"name"`
	Players []Player `json:"players"`
	DateCreated time.Time `json:"dateCreated,omitempty"`
	DateModified time.Time `json:"dateModified,omitempty"`
}
