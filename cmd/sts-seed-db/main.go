package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/heyztb/steam-trade-server/internal/database"
)

func main() {
	file, err := os.Open("accounts.json")
	if err != nil {
		panic(err)
	}

	accounts, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var bots []*database.Bot
	err = json.Unmarshal(accounts, &bots)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", "steam-trade-server.db")
	if err != nil {
		panic(err)
	}

	queries := database.New(db)

	for _, bot := range bots {
		err = queries.InsertBot(context.Background(), database.InsertBotParams{
			Username:       bot.Username,
			Passwd:         bot.Passwd,
			SharedSecret:   bot.SharedSecret,
			IdentitySecret: bot.IdentitySecret,
		})
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Database loaded with accounts")
}
