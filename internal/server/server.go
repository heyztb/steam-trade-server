package server

import (
	"context"
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/heyztb/steam-trade-server/internal/router"
)

type steamTradeServer struct {
	closers []io.Closer
	server  *http.Server
}

func (s *steamTradeServer) Start() error {
	go func() {
		log.Default().Println("Server listening on 127.0.0.1:3000")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Fatal server error: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	log.Default().Println("Server shutting down")

	for _, c := range s.closers {
		c.Close()
	}

	return s.server.Shutdown(ctx)
}

func Setup() *steamTradeServer {
	db, err := sql.Open("sqlite3", "steam-trade-server.db")
	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		return nil
	}

	router := router.NewRouter(db)
	httpServer := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	return &steamTradeServer{
		server:  httpServer,
		closers: []io.Closer{db},
	}
}
