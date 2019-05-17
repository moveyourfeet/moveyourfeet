package main

import (
	"net/http"

	customHttp "github.com/moveyourfeet/gameManager/http"

	"github.com/gorilla/mux"
	"github.com/moveyourfeet/gameManager/game"
	"github.com/moveyourfeet/gameManager/middleware"
	customRouter "github.com/moveyourfeet/gameManager/router"
)

// NewRouter sets up all http endpoints
func NewRouter() *mux.Router {

	//init router
	router := mux.NewRouter()

	customRouter.AppRoutes = append(customRouter.AppRoutes, game.Routes, customHttp.Routes)

	for _, route := range customRouter.AppRoutes {

		//create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//loop through each sub route
		for _, r := range route.SubRoutes {

			var handler http.Handler
			handler = r.HandlerFunc

			//check to see if route should be protected with jwt
			if r.Protected {
				handler = middleware.JWTMiddleware(r.HandlerFunc)
			}

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
