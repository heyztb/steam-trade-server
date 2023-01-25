package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/heyztb/steam-trade-server/internal/handlers"
	_ "github.com/joho/godotenv/autoload"
)

// NewRouter creates a new router for the spec and the given handlers.
// Steam Trade Server
//
// # RESTful API to enable users to trade Steam items to and from automated trading accounts
//
// 0.1.0
func NewRouter() http.Handler {
	postgresDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", postgresDSN)
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
