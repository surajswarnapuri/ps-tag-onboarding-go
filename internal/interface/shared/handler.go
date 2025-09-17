// Package shared contains shared code for the api layer.
package shared

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is resonsible for defining a HTTP request route and corresponding handler function.
type Handler struct {
	Route func(r *mux.Route)
	Func  http.HandlerFunc
}

func (h Handler) AddRoute(r *mux.Router) {
	h.Route(r.NewRoute().HandlerFunc(h.Func))
}
