package game

import "github.com/georace/game-manager/router"

var Routes = router.RoutePrefix{
	"/games",
	[]router.Route{
		router.Route{"GamesIndex", "GET", "", IndexHandler, false},
		router.Route{"GamesShow", "GET", "/{userId}", ShowHandler, true},
		router.Route{"GamesCreate", "POST", "", CreateHandler, false},
		router.Route{"DeleteHandler", "DELETE", "/{userId}", DeleteHandler, true},
		router.Route{"UpdateHandler", "PUT", "/{userId}", UpdateHandler, true},
	},
}
