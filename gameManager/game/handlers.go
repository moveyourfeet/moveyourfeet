package game

import (
	"encoding/json"
	"fmt"
	"net/http"

	customHTTP "github.com/georace/gameManager/http"
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
// @Tags Games
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	games, err := List()
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))

	}
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
// @Tags Games
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	game, err := Get(params["gameId"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, fmt.Sprintf("%v", err))
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
// @Tags Games
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var newGame NewGame
	err := json.NewDecoder(r.Body).Decode(&newGame)
	game, err := Create(newGame)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}
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
// @Tags Games
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	games, err := Delete(params["gameId"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadGateway, "Error: "+err.Error())
		return
	}

	customHTTP.NewResponse(w, games)
}
