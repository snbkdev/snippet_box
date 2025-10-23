package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr string
	staticDir string
}

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4060", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		//AddSource: true,
	}))

	app := &application{
		logger: logger,
	}

	logger.Info("strating server on", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}