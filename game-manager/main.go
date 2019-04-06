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
		panic("No database connection")
	}
	// http://doc.gorm.io/database.html#migration
	db.DB.DropTableIfExists("games")
	db.SetupProcedures()

	db.DB.AutoMigrate(&user.User{})
	db.DB.AutoMigrate(&game.Game{})

	db.SetupTriggers()

	defer db.DB.Close()

	//create http server
	log.Fatal(http.ListenAndServe(":"+port, router))
}
