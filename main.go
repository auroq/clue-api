package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

type Player struct {
	ID string `json:"id,omitempty"`
	human bool `json:"human,omitempty"`
	name string `json:"name,omitempty"`
	dateCreated time.Time `json:"dateCreated,omitempty"`
	dateModified time.Time `json:"dateModified,omitempty"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPlayers).Methods("GET")
	router.HandleFunc("/people/{id}", GetPlayer).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePlayer).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePlayer).Methods("DELETE")

	connectToMongoDB()
	fmt.Println("Starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

var players []Player

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, player := range players {
		if player.ID == params["id"] {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	json.NewEncoder(w).Encode(&Player{})
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	_ = json.NewDecoder(r.Body).Decode(&player)
	player.ID = params["id"]
	player.dateCreated = time.Now()
	player.dateModified = time.Now()
	players = append(players, player)
	json.NewEncoder(w).Encode(players)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, player := range players {
		if player.ID == params["ID"] {
			players = append(players[:i], players[i+1:]...)
			break
		}
		json.NewEncoder(w).Encode(players)
	}
}
