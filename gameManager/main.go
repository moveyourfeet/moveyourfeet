package main

import (
	"log"
	"net/http"
	"os"

	"github.com/georace/game-manager/db"
	"github.com/georace/game-manager/game"
	"github.com/subosito/gotenv"

	_ "github.com/georace/game-manager/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Game Manager
// @version 1.0
// @description Handles games, game rules, and join codes.

// @host
// @BasePath /
func main() {
	gotenv.Load()
	//init router
	port := os.Getenv("PORT")
	router := NewRouter()
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	//Setup database
	db.DB = db.SetupDB()
	if db.DB == nil {
		panic("No database connection")
	}
	// http://doc.gorm.io/database.html#migration
	// db.DB.DropTableIfExists("games")
	db.SetupProcedures()

	db.DB.AutoMigrate(&game.Game{})

	db.SetupTriggers()

	defer db.DB.Close()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}
