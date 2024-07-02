package app

import "net/http"

// * Decoupling:
// Route is an http.handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

func NewServeMux(route1, route2 Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(route1.Pattern(), route1)
	mux.Handle(route2.Pattern(), route2)

	return mux
}
