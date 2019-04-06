package main

import (
	"log"
	"net/http"
	"os"

	"github.com/georace/game-manager/game"
	"github.com/georace/game-manager/user"

	"github.com/georace/game-manager/db"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	//init router
	port := os.Getenv("PORT")
	router := NewRouter()

	//Setup database
	db.DB = db.SetupDB()
	if db.DB == nil {
		panic("ARGHHGH")
	}
	// http://doc.gorm.io/database.html#migration
	db.DB.AutoMigrate(&user.User{})
	db.DB.AutoMigrate(&game.Game{})
	defer db.DB.Close()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}
