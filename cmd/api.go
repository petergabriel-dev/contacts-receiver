package main

import (
	"contacts/internal/contact"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

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

	contactService := contact.NewService()
	contactHandler := contact.NewHandler(contactService)
	r.Get("/contact", contactHandler.ListContacts)

	return r
}

func (app *application) run(h http.Handler) error {
	// configure the server
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Starting server on port ", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string // 8080
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
