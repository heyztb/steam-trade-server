package models

import (
	"log"
	"time"

	"github.com/Philipp15b/go-steam/v3"
	"github.com/Philipp15b/go-steam/v3/totp"
	"github.com/Philipp15b/go-steam/v3/tradeoffer"
	"github.com/heyztb/steam-trade-server/internal/database"
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
		case *steam.LoggedOnEvent:
			bot.Logger.Printf("Log on succeeded: %s", e.ClientSteamId.String())
		case *steam.LogOnFailedEvent:
			bot.Logger.Printf("Log on failed: %+v\n", e.Result)
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

func (bot *TradeBot) Trade(tradeUrl string, ourItems []Asset, theirItems []Asset) {
	_ = tradeoffer.NewClient("", bot.Client.Web.SessionId, bot.Client.Web.SteamLogin, bot.Client.Web.SteamLoginSecure)

	// TODO: Implement
}
