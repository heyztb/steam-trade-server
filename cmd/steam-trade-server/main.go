package main

import (
	"github.com/heyztb/steam-trade-server/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := server.Setup().Start()
	if err != nil {
		panic(err)
	}
}
