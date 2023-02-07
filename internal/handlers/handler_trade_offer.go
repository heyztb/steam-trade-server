package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Philipp15b/go-steam/v3"
	"github.com/heyztb/steam-trade-server/internal/database"
	"github.com/heyztb/steam-trade-server/internal/models"
	"github.com/pkg/errors"
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
	w.Header().Set("Content-Type", "application/json")

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
		return
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

		handler.logger.Printf("Error creating trade offer: %s\n", err.Error())
		return
	}

	w.WriteHeader(201)
	w.Header().Set("Location", tradeOffer)
	w.Write(nil)
}

func (handler *createTradeOfferHandler) createTradeOffer(proposal *models.TradeProposal) (string, error) {
	queries := database.New(handler.db)

	err := proposal.Validate()
	if err != nil {
		return "", errors.Wrap(err, "failed to validate proposal")
	}

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
		botCreds, err := queries.GetRandomBot(context.Background())
		if err != nil {
			return "", errors.Wrap(err, "failed to get bot from database")
		}

		tradeBot := &models.TradeBot{
			Credentials: &botCreds,
			Client:      steam.NewClient(),
			Logger:      log.Default(),
		}

		go tradeBot.EventHandler()

		tradeBot.Login()
		offerId, err := tradeBot.Trade(proposal.TradeUrl, proposal.Offer, proposal.Want)
		if err != nil {
			return "", errors.Wrap(err, "failed to create trade offer")
		}

		return fmt.Sprintf("https://steamcommunity.com/tradeoffer/%d", *offerId), nil
	case ToUser:
		return "", errors.New("not implemented")
	case Mutual:
		return "", errors.New("not implemented")
	}

	return "", nil
}
