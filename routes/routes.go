package routes

import (
	"SolarInstaller/routes/authRoutes"
	"SolarInstaller/routes/projectRoutes"
	"SolarInstaller/routes/todoRoutes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// InitRoutes initializes all module routes
func InitRoutes() http.Handler {
	r := chi.NewRouter()

	// Mount routes for todos, projects, and auth
	r.Mount("/todos", todoRoutes.TodoRouter())
	r.Mount("/projects", projectRoutes.ProjectRouter())
	r.Mount("/auth", authRoutes.AuthRouter())

	return r
}
