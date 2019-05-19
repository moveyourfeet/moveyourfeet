package game

import (
	"errors"
	"time"
)

// Game represents a user created game
type Game struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Name       string    `json:"name"`
	Owner      string    `json:"owner"`
	CreateTime time.Time `json:"create_time"`
	JoinSecret string    `json:"join_secret"`
}

// NewGame is a request from a user to create a game
type NewGame struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func validateNewGame(newGame NewGame) (Game, error) {

	var game Game

	if newGame.Name == "" {
		return game, errors.New("missing field 'name' in JSON object")
	}

	if newGame.Owner == "" {
		return game, errors.New("missing field 'owner' in JSON object")
	}

	game.Name = newGame.Name
	game.Owner = newGame.Owner
	game.CreateTime = time.Now().UTC()
	return game, nil
}
