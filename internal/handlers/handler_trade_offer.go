package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/heyztb/steam-trade-server/internal/models"
)

type tradeDirection string

const (
	ToServer tradeDirection = "ToServer"
	ToUser   tradeDirection = "ToUser"
	Mutual   tradeDirection = "Mutual"
)

type createTradeOfferHandler struct {
	db     *sql.DB
	logger *log.Logger
}

func NewCreateTradeOfferHandler(db *sql.DB, logger *log.Logger) *createTradeOfferHandler {
	return &createTradeOfferHandler{
		db:     db,
		logger: logger,
	}
}

func (handler *createTradeOfferHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tp := &models.TradeProposal{}

	if err := json.NewDecoder(r.Body).Decode(tp); err != nil {
		errMsg := models.Error{
			Code:    400,
			Message: fmt.Sprintf("Invalid request body: %s", err.Error()),
		}
		errJson, _ := json.Marshal(errMsg)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(errJson)

		handler.logger.Printf("Error parsing request body: %s\n", err.Error())
	}

	// TODO: Add logic
	// var direction tradeDirection
	// switch {
	// case len(tp.Offer) == 0 && len(tp.Want) >= 1:
	// 	direction = ToServer
	// case len(tp.Offer) >= 1 && len(tp.Want) == 0:
	// 	direction = ToUser
	// default:
	// 	direction = Mutual
	// }

	// if direction == ToServer {
	// }

	// if direction == ToUser {
	// }

	// if direction == Mutual {
	// }
}
