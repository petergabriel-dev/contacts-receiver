package main

import (
	repo "contacts/internal/adapters/sqlite/sqlc"
	"contacts/internal/contact"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	db     *sql.DB
}

// Mount the server
// Setup the routes and middleware
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID) //Rate-Limiting
	r.Use(middleware.RealIP)    // Rate-Limiting and Analytics and Tracing
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running"))
	})

	// GET All Contacts
	contactService := contact.NewService(repo.New(app.db))
	contactHandler := contact.NewHandler(contactService)
	r.Get("/contacts", contactHandler.ListContacts)
	// GET Contact by ID
	r.Get("/contacts/{id}", contactHandler.GetContactByID)

	return r
}

// Run the server
// Server Configurations
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Starting server on port", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
