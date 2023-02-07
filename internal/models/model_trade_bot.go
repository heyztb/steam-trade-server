package models

import (
	"log"
	"time"

	"github.com/Philipp15b/go-steam/v3"
	"github.com/Philipp15b/go-steam/v3/steamid"
	"github.com/Philipp15b/go-steam/v3/totp"
	"github.com/Philipp15b/go-steam/v3/trade/tradeapi"
	"github.com/Philipp15b/go-steam/v3/tradeoffer"
	"github.com/heyztb/steam-trade-server/internal/database"
	"github.com/heyztb/steam-trade-server/pkg/tradeurl"
	"github.com/pkg/errors"
)

type TradeBot struct {
	Credentials *database.Bot
	Client      *steam.Client
	Logger      *log.Logger
}

func (bot *TradeBot) EventHandler() {
	for event := range bot.Client.Events() {
		switch e := event.(type) {
		case *steam.FatalErrorEvent:
			bot.Logger.Printf("Fatal error")
			return
		case *steam.LoggedOnEvent:
			bot.Logger.Printf("Log on succeeded: %s", e.ClientSteamId.String())
		case *steam.LogOnFailedEvent:
			bot.Logger.Printf("Log on failed: %+v\n", e.Result)
			return
		case *steam.WebSessionIdEvent:
			bot.Client.Web.LogOn()
		case *steam.TradeResultEvent:
			switch e.Response {
			case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 22, 32, 33:
				bot.Logger.Printf("Error with trade offer: %d", e.Response)
				return
			case 0, 50:
				bot.Logger.Printf("Good trade: %d", e.Response)
				return
			}
		}
	}
}

func (bot *TradeBot) Login() error {
	twoFactorCode, err := totp.GenerateTotpCode(bot.Credentials.SharedSecret, time.Now())
	if err != nil {
		return errors.Wrap(err, "Failed to generate TOTP code")
	}

	_, err = bot.Client.Connect()
	if err != nil {
		return errors.Wrap(err, "Failed to connect to Steam")
	}

	bot.Client.Auth.LogOn(&steam.LogOnDetails{
		Username:      bot.Credentials.Username,
		Password:      bot.Credentials.Passwd,
		TwoFactorCode: twoFactorCode,
	})

	return nil
}

func (bot *TradeBot) Trade(tradeUrl string, ourItems []Asset, theirItems []Asset) (*uint64, error) {
	tradeClient := tradeoffer.NewClient(tradeoffer.APIKey(bot.Credentials.ApiKey), bot.Client.Web.SessionId, bot.Client.Web.SteamLogin, bot.Client.Web.SteamLoginSecure)

	tradeInfo, err := tradeurl.Parse(tradeUrl)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse trade url")
	}

	partnerId := tradeInfo[0]
	accessToken := tradeInfo[1]

	theirId, err := steamid.NewId(partnerId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse steam ID")
	}

	us := []tradeoffer.TradeItem{}
	them := []tradeoffer.TradeItem{}

	for _, item := range ourItems {
		us = append(us, tradeoffer.TradeItem{
			AppId:      uint32(item.AppID),
			ContextId:  2,
			Amount:     1,
			AssetId:    uint64(item.AssetID),
			CurrencyId: 0,
		})
	}

	for _, item := range theirItems {
		them = append(them, tradeoffer.TradeItem{
			AppId:      uint32(item.AppID),
			ContextId:  2,
			Amount:     1,
			AssetId:    uint64(item.AssetID),
			CurrencyId: 0,
		})
	}

	offerId, err := tradeClient.Create(theirId, &accessToken, us, them, nil, "")
	if err != nil {
		return nil, errors.Wrap(err, "unable to create trade offer")
	}

	return &offerId, nil
}

func (bot *TradeBot) Confirm(other steamid.SteamId) error {
	tradeApiClient := tradeapi.New(bot.Client.Web.SessionId, bot.Client.Web.SteamLogin, bot.Client.Web.SteamLoginSecure, other)

	_, err := tradeApiClient.SetReady(true)
	if err != nil {
		return errors.Wrapf(err, "unable to set ready on trade w/ %s", other.String())
	}

	_, err = tradeApiClient.Confirm()
	if err != nil {
		return errors.Wrapf(err, "unable to confirm trade w/ %s", other.String())
	}

	return nil
}
