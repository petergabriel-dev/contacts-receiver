package main

import (
	"contacts/internal/env"
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "./internal/adapters/sqlite/database/database"),
		},
	}

	// Structure Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database Connection
	conn, er := sql.Open("sqlite3", cfg.db.dsn)
	if er != nil {
		slog.Error("Failed to connect to database", "error", er)
		os.Exit(1)
	}
	defer conn.Close()

	logger.Info("Database connection established", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start ", "error", err)
		os.Exit(1)
	}
}
