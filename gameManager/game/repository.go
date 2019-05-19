package game

import (
	"fmt"

	"github.com/moveyourfeet/gameManager/db"
)

// List gets all games
func List() ([]Game, error) {
	var games []Game
	//since we're passing a pointer to games, db.Find assigns array to the address
	err := db.DB.Find(&games).Error
	if err != nil {
		return games, fmt.Errorf("Error listing games: %v", err)
	}
	return games, nil
}

// Get Gets a game with id
func Get(gameID string) (Game, error) {

	var game Game
	if db.DB.First(&game, gameID).RecordNotFound() {
		return game, fmt.Errorf("No record found with id: %s", gameID)
	}
	return game, nil
}

// Create Creates game
func Create(newGame NewGame) (Game, error) {
	var game Game
	game, err := validateNewGame(newGame)
	if err != nil {
		return game, err
	}

	err = db.DB.Create(&game).Error
	if err != nil {
		return game, err
	}
	// Reload from database to get the join_secret from the trigger
	db.DB.First(&game, game.ID)
	return game, nil
}

// Delete game by id
func Delete(gameID string) ([]Game, error) {
	var game Game
	var games []Game

	db.DB.First(&game, gameID)
	db.DB.Delete(&game)

	db.DB.Find(&games)
	return games, nil
}
