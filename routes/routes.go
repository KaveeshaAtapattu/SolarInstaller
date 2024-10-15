package routes

import (
	"SolarInstaller/routes/projectRoutes"
	"SolarInstaller/routes/todoRoutes"
	"net/http"

	"github.com/go-chi/chi"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()

	// Mount specific routes for each module
	r.Mount("/todos", todoRoutes.TodoRouter())
	r.Mount("/projects", projectRoutes.ProjectRouter())

	return r
}
