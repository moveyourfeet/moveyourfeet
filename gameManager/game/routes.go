package game

import "github.com/moveyourfeet/gameManager/router"

// Routes contains all endpoints defined in this package
var Routes = router.RoutePrefix{
	"/games",
	[]router.Route{
		router.Route{"GamesIndex", "GET", "", IndexHandler, false},
		router.Route{"GamesShow", "GET", "/{gameId}", ShowHandler, false},
		router.Route{"GamesCreate", "POST", "", CreateHandler, false},
		router.Route{"DeleteHandler", "DELETE", "/{gameId}", DeleteHandler, true},
		// router.Route{"UpdateHandler", "PUT", "/{gameId}", UpdateHandler, true},
	},
}
