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
	w.Header().Add("Content-Type", "application/json")

	tp := &models.TradeProposal{}

	if err := json.NewDecoder(r.Body).Decode(tp); err != nil {
		errMsg := models.Error{
			Code:    400,
			Message: fmt.Sprintf("Invalid request body: %s", err.Error()),
		}
		errJson, _ := json.Marshal(errMsg)

		w.WriteHeader(400)
		w.Write(errJson)

		handler.logger.Printf("Error parsing request body: %s\n", err.Error())
	}

	tradeOffer, err := handler.createTradeOffer(tp)
	if err != nil {
		errMsg := models.Error{
			Code:    400,
			Message: fmt.Sprintf("Error creating trade offer: %s", err.Error()),
		}
		errJson, _ := json.Marshal(errMsg)

		w.WriteHeader(400)
		w.Write(errJson)

		handler.logger.Printf("Error parsing request body: %s\n", err.Error())
	}

	w.WriteHeader(201)
	w.Header().Add("Location", tradeOffer)
	w.Write([]byte(`{}`))
}

// TODO
func (h *createTradeOfferHandler) createTradeOffer(proposal *models.TradeProposal) (string, error) {

	var direction tradeDirection
	switch {
	case len(proposal.Offer) == 0 && len(proposal.Want) > 0:
		direction = ToServer
	case len(proposal.Offer) > 0 && len(proposal.Want) == 0:
		direction = ToUser
	default:
		direction = Mutual
	}

	switch direction {
	case ToServer:
		panic("not implemented")
	case ToUser:
		panic("not implemented")
	case Mutual:
		panic("not implemented")
	}

	return "", nil
}
