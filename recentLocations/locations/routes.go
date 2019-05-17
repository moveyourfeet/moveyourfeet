package locations

import "github.com/moveyourfeet/moveyourfeet/recentLocations/router"

var Routes = router.RoutePrefix{
	"/locations",
	[]router.Route{
		router.Route{"Show", "GET", "/{gameId}", ShowHandler, false},
	},
}
