package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bruno-holanda15/api-rest-challenge-rocketseat/api"
)

func main() {
	handler := api.NewHTTPHandler()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	server := http.Server{
		Addr:                         ":8088",
		Handler:                      handler, // multiplexer | mux
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  time.Second * 5,
		WriteTimeout:                 time.Second * 5,
		ErrorLog:                     slog.NewLogLogger(log.Handler(), slog.LevelDebug),
		IdleTimeout:                  time.Minute * 2,
	}

	errChan := make(chan os.Signal, 1)
	signal.Notify(errChan, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting HTTP Server")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	<-errChan
	log.Info("server is shutting down")
}
