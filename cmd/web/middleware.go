package main

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/helpers"
)

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated
		if !helpers.IsAuthenticated(r) {
			// Store the error message in the session
			session.Put(r.Context(), "error", "Log In First")
			// Redirect to the login page
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}
		// If authenticated, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
