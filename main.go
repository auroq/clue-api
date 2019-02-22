package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPlayers).Methods("GET")
	router.HandleFunc("/people/{id}", GetPlayer).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePlayer).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePlayer).Methods("DELETE")
	fmt.Println("Starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {}
func GetPlayer(w http.ResponseWriter, r *http.Request) {}
func CreatePlayer(w http.ResponseWriter, r *http.Request) {}
func DeletePlayer(w http.ResponseWriter, r *http.Request) {}

type Player struct {
	id string `json:"id,omitempty"`
	human bool `json:"human,omitempty"`
	name string `json:"name,omitempty"`
	dateCreated time.Time `json:"dateCreated,omitempty"`
	dateModified time.Time `json:"dateModified,omitempty"`
}