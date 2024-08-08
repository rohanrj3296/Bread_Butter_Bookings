package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rohanrj3296/Bread_Butter_Bookings/pkg/config"
	"github.com/rohanrj3296/Bread_Butter_Bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability",handlers.Repo.Availability)
	//below line catches any request that is posted to this url and takes it to right handler
	mux.Post("/search-availability",handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json",handlers.Repo.AvailabilityJSON)
	mux.Get("/contact",handlers.Repo.Contact)
	mux.Get("/make-reservation",handlers.Repo.Reservation)
	//creating a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	//starting the file server
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
