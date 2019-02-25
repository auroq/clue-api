package main

import (
	"fmt"
	"github.com/auroq/clue-api/pkg/data"
	"github.com/auroq/clue-api/pkg/game"
	"github.com/auroq/clue-api/pkg/player"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"log"
	"net/http"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(data.NewConfiguration)
	container.Provide(data.NewDbConnection)
	container.Provide(game.NewGameService)
	container.Provide(player.NewPlayerService)
	container.Provide(game.NewGameController)
	return container
}

func main() {
	container := BuildContainer()

	err := container.Invoke(func(gameController game.Controller){
		router := mux.NewRouter()
		gameController.Endpoints(router)
		fmt.Println("Starting on port 8000")
		log.Fatal(http.ListenAndServe(":8000", router))
	})
	if err != nil {
		log.Fatal(err)
	}
}

