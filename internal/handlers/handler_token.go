package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/heyztb/steam-trade-server/internal/models"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const (
	audience = `http://localhost:3000`
	issuer   = `https://github.com/heyztb/steam-trade-server`
)

type tokenHandler struct {
	logger *log.Logger
}

func NewTokenHandler(logger *log.Logger) *tokenHandler {
	return &tokenHandler{
		logger: logger,
	}
}

func (handler *tokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sessionId := uuid.New()

	jwt.Settings(jwt.WithFlattenAudience(true))

	tok, err := jwt.NewBuilder().
		Audience([]string{audience}).
		Issuer(issuer).
		NotBefore(time.Now()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(3600*time.Second)).
		Claim("session_id", sessionId.String()).
		Build()

	if err != nil {
		handler.logger.Printf("error building token: %s\n", err.Error())
		errMsg := &models.Error{
			Code:    500,
			Message: fmt.Sprintf("Error building token: %s", err.Error()),
		}
		errJson, _ := json.Marshal(errMsg)

		w.WriteHeader(500)
		w.Write(errJson)
		return
	}

	key, err := jwk.FromRaw([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		handler.logger.Printf("error fetching key: %s\n", err.Error())
		errMsg := &models.Error{
			Code:    500,
			Message: fmt.Sprintf("Error fetching key: %s", err.Error()),
		}
		errJson, _ := json.Marshal(errMsg)

		w.WriteHeader(500)
		w.Write(errJson)
		return
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, key))
	if err != nil {
		handler.logger.Printf("error signing jwt: %s\n", err.Error())
		errMsg := &models.Error{
			Code:    500,
			Message: fmt.Sprintf("Error signing JWT: %s", err.Error()),
		}
		errJson, _ := json.Marshal(errMsg)

		w.WriteHeader(500)
		w.Write(errJson)
		return
	}

	token := &models.AuthToken{
		Token:     string(signed),
		TokenType: "Bearer",
		ExpiresIn: 3600,
	}

	json, _ := json.Marshal(token)

	w.WriteHeader(200)
	w.Write(json)
}
