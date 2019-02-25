package data

import "time"

var client dataStore

func getClient() dataStore {
	if client == nil {
		client, _ = getConnection()
	}

	return client
}

func addPlayer(name string, human bool) (player Player, err error) {
	player = Player {
		Name: name,
		Human: human,
		DateCreated: time.Now(),
		DateModified: time.Now(),
	}
	player.ID, err = getClient().insert("clue-api", "players", player)
	return player, err
}

func CreateGame(name string, playerNames []string) (game Game, err error) {
	game.Name = name
	for _, playerName := range playerNames {
		player, err := addPlayer(playerName, true)
		if err != nil {
			return game, err
		}
		game.Players = append(game.Players, player)
	}

	game.ID, err = getClient().insert("clue-api", "games", game)
	return game, err
}