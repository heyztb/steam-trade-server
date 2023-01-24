package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter creates a new router for the spec and the given handlers.
// Steam Trade Server
//
// # RESTful API to enable users to trade Steam items to and from automated trading accounts
//
// 0.1.0
func NewRouter() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	return r
}
