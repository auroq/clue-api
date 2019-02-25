package main

import (
	"fmt"
	"github.com/auroq/clue-api/pkg/game"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	game.Endpoints(router)
	fmt.Println("Starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

