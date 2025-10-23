package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"flag"
	"log"
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

	connStr := "user=user password=password dbname=dbname host=host port=port sslmode=disable"
    db, err := openDB(connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	app := &application{
		logger: logger,
	}

	logger.Info("strating server on", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}