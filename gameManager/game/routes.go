package game

import "github.com/moveyourfeet/gameManager/router"

// Routes contains all endpoints defined in this package
var Routes = router.RoutePrefix{
	Prefix: "/games",
	SubRoutes: []router.Route{
		{
			Name:        "GamesIndex",
			Method:      "GET",
			Pattern:     "",
			HandlerFunc: IndexHandler,
			Protected:   false},
		{
			Name:        "GamesShow",
			Method:      "GET",
			Pattern:     "/{gameId}",
			HandlerFunc: ShowHandler,
			Protected:   false},
		{
			Name:        "GamesCreate",
			Method:      "POST",
			Pattern:     "",
			HandlerFunc: CreateHandler,
			Protected:   false},
		{
			Name:        "DeleteHandler",
			Method:      "DELETE",
			Pattern:     "/{gameId}",
			HandlerFunc: DeleteHandler,
			Protected:   true},
	},
}
