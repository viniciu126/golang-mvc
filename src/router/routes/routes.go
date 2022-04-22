package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes represents all API routes
type Routes struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

// Configure configure all routes inside a router
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		routeFunction := middlewares.Logger(route.Function)

		if route.Auth {
			routeFunction = middlewares.Auth(routeFunction)
		}

		r.HandleFunc(route.URI, routeFunction).Methods(route.Method)
	}

	return r
}
