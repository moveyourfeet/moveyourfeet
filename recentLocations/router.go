package main

import (
	"net/http"

	"github.com/georace/recentLocations/locations"
	customRouter "github.com/georace/recentLocations/router"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	customRouter.AppRoutes = append(customRouter.AppRoutes, locations.Routes)

	for _, route := range customRouter.AppRoutes {

		//create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//loop through each sub route
		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			//attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}
	}
	return router
}
