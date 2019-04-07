package game

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/georace/game-manager/db"
	customHTTP "github.com/georace/game-manager/http"
	"github.com/gorilla/mux"
)

// IndexHandler godoc
// @Summary List games
// @Description get games
// @Accept  json
// @Produce  json
// @Success 200 {array} game.Game
// @Failure 404 {object} http.ErrorResponse
// @Router /games [get]
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var games []Game
	//since we're passing a pointer to games, db.Find assigns array to the address
	db.DB.Find(&games)
	customHTTP.NewResponse(w, games)
}

// ShowHandler godoc
// @Summary Get game
// @Description get one game
// @Accept  json
// @Produce  json
// @Success 200 {object} game.Game
// @Failure 404 {object} http.ErrorResponse
// @Param gameId path int true "Game ID"
// @Router /games/{gameId} [get]
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game Game
	fmt.Print(params)
	if db.DB.First(&game, params["gameId"]).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "No record found with id: "+params["gameId"])
		return
	}
	customHTTP.NewResponse(w, game)
}

// CreateHandler godoc
// @Summary Create game
// @Description create a game
// @Accept  json
// @Produce  json
// @Param newgame body game.NewGame true "Add game"
// @Success 200 {object} game.Game
// @Failure 404 {object} http.ErrorResponse
// @Router /games [post]
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var newGame NewGame
	err := json.NewDecoder(r.Body).Decode(&newGame)

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	game, err := validateNewGame(newGame)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	err = db.DB.Create(&game).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	// Reload from database to get the join_secret from the trigger
	db.DB.First(&game, game.ID)
	customHTTP.NewResponse(w, game)
}

// DeleteHandler godoc
// @Summary Delete game
// @Description deletes a game
// @Accept  json
// @Produce  json
// @Success 200 {object} game.Game
// @Failure 404 {object} http.ErrorResponse
// @Router /games/{gameId} [delete]
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game Game
	var games []Game

	db.DB.First(&game, params["gameId"])
	db.DB.Delete(&game)

	db.DB.Find(&games)
	customHTTP.NewResponse(w, games)
}

// UpdateHandler godoc
// @Summary Delete game
// @Description deletes a game
// @Accept  json
// @Produce  json
// @Success 200 {object} game.Game
// @Failure 404 {object} http.ErrorResponse
// @Router /games/{gameId} [put]
// func UpdateHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var game Game
// 	reqgameId := r.Header.Get("gameId")

// 	w.Header().Set("Content-Type", "application/json")
// 	if params["gameId"] != reqgameId {
// 		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Not allowed to edit other games")
// 		return
// 	}
// 	db.DB.First(&game, params["gameId"])
// 	db.DB.Model(&game).Update("name", r.FormValue("name"))
// 	json.NewEncoder(w).Encode(&game)
// }
