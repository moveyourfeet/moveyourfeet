package router

import "net/http"

// AppRoutes contains all routes for this service
var AppRoutes []RoutePrefix

// RoutePrefix allows for defining a prefix for all of the sub routes
type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

// Route describes a route, method and a handler for the route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}
