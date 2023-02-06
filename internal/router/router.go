package router

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/jwtauth/v5"
	"github.com/heyztb/steam-trade-server/internal/handlers"
	"github.com/heyztb/steam-trade-server/internal/models"
)

// NewRouter creates a new router for the spec and the given handlers.
// Steam Trade Server
//
// # RESTful API to enable users to trade Steam items to and from automated trading accounts
//
// 0.1.0

func NewRouter(db *sql.DB) http.Handler {
	logger := log.Default()
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(httprate.Limit(
			1,
			time.Hour,
			httprate.WithKeyByRealIP(),
			httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
				errMsg := models.Error{
					Code:    429,
					Message: "Too many requests",
				}
				errJson, _ := json.Marshal(errMsg)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(429)
				w.Write(errJson)
			}),
		))
		r.Method("GET", "/token", handlers.NewTokenHandler(logger))
	})

	r.Group(func(r chi.Router) {
		tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(httprate.Limit(
			1,
			time.Second,
			httprate.WithKeyFuncs(httprate.KeyByRealIP, func(r *http.Request) (string, error) {
				token := r.Context().Value(jwtauth.TokenCtxKey).(jwt.Token)
				sessionId := token.PrivateClaims()["session_id"].(string)
				return sessionId, nil
			}),
			httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
				errMsg := models.Error{
					Code:    429,
					Message: "Too many requests",
				}
				errJson, _ := json.Marshal(errMsg)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(429)
				w.Write(errJson)
			}),
		))

		r.Method("POST", "/trade/new", handlers.NewCreateTradeOfferHandler(db, logger))
	})

	return r
}
