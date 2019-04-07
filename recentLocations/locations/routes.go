package locations

import "github.com/georace/recentLocations/router"

var Routes = router.RoutePrefix{
	"/locations",
	[]router.Route{
		router.Route{"Show", "GET", "/{gameId}", ShowHandler, false},
	},
}
