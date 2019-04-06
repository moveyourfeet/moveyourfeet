package game

import (
	"encoding/json"
	"net/http"

	"github.com/georace/game-manager/db"
	customHTTP "github.com/georace/game-manager/http"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var games []Game
	//since we're passing a pointer to games, db.Find assigns array to the address
	db.DB.Find(&games)
	customHTTP.NewResponse(w, games)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game Game
	db.DB.First(&game, params["gameId"])
	customHTTP.NewResponse(w, game)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var game Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	err = db.DB.Create(&game).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	customHTTP.NewResponse(w, game)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game Game
	var games []Game

	db.DB.First(&game, params["gameId"])
	db.DB.Delete(&game)

	db.DB.Find(&games)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var game Game
	reqgameId := r.Header.Get("gameId")

	w.Header().Set("Content-Type", "application/json")
	if params["gameId"] != reqgameId {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Not allowed to edit other games")
		return
	}
	db.DB.First(&game, params["gameId"])
	db.DB.Model(&game).Update("name", r.FormValue("name"))
	json.NewEncoder(w).Encode(&game)
}
