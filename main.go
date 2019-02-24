package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var client dataStore

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/players", GetPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", GetPlayer).Methods("GET")
	router.HandleFunc("/players", CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", DeletePlayer).Methods("DELETE")

	client, _ = getConnection()

	fmt.Println("Starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

var players []Player

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//for _, player := range players {
	//	if player.ID == params["id"] {
	//		json.NewEncoder(w).Encode(player)
	//		return
	//	}
	//}
	//json.NewEncoder(w).Encode(&Player{})
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode("Invalid request")
		return
	}
	player.DateCreated = time.Now()
	player.DateModified = time.Now()
	players = append(players, player)
	player.ID, err = client.insert("clue-api", "players", player)
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(player)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//for i, player := range players {
	//	if player.ID == params["ID"] {
	//		players = append(players[:i], players[i+1:]...)
	//		break
	//	}
	//	json.NewEncoder(w).Encode(players)
	//}
}
