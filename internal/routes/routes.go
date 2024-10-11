// Package routes provides routes setup function
package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zongrade/testeffective/internal/handlers"
)

// SetupRoutes setup base routes.
func SetupRoutes() {
	http.HandleFunc("/songs", handlers.SongsHandler)
	http.HandleFunc("/info", handlers.InfoHandler)
	http.HandleFunc("/couplets", handlers.CoupletsHandler)
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}
