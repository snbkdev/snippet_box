package main

import (
	"database/sql"
	"flag"
	"lets_go/internal/models"
	"log/slog"
	"net/http"
	"os"
	"html/template"

	_ "github.com/lib/pq"
)

type config struct {
	addr string
	staticDir string
}

type application struct {
	logger *slog.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4060", "HTTP network address")
	connStr := "user=user password=password dbname=dbname host=host port=port sslmode=disable"
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := openDB(connStr)
    if err != nil {
        logger.Error(err.Error())
		os.Exit(1)
    }
    defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
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