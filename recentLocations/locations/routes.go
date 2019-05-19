package locations

import "github.com/moveyourfeet/moveyourfeet/recentLocations/router"

// Routes from the /locations/.. package
var Routes = router.RoutePrefix{
	Prefix: "/locations",
	SubRoutes: []router.Route{
		{
			Name:        "Show",
			Method:      "GET",
			Pattern:     "/{gameId}",
			HandlerFunc: ShowHandler,
			Protected:   false},
	},
}
