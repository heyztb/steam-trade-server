package router

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/heyztb/steam-trade-server/internal/handlers"
)

// NewRouter creates a new router for the spec and the given handlers.
// Steam Trade Server
//
// # RESTful API to enable users to trade Steam items to and from automated trading accounts
//
// 0.1.0

func NewRouter() http.Handler {

	db, err := sql.Open("sqlite3", "steam-trade-server.db")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Method("POST", "/trade/new", handlers.NewCreateTradeOfferHandler(db, log.Default()))

	return r
}
