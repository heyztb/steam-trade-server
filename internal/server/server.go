package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/heyztb/steam-trade-server/internal/router"
)

func Start() {
	router := router.NewRouter()
	httpServer := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		log.Default().Println("Server listening on 127.0.0.1:3000")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Fatal server error: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	httpServer.Shutdown(ctx)
	log.Default().Println("Server shutting down")
}
