package authRoutes

import (
	"SolarInstaller/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// AuthRouter handles authentication routes
func AuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/register", controllers.RegisterHandler)
	r.Post("/login", controllers.LoginHandler)
	return r
}
