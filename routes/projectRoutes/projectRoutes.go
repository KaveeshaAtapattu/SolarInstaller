package projectRoutes

import (
	"SolarInstaller/controllers"
	"net/http"

	"github.com/go-chi/chi"
)

func ProjectRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/", controllers.CreateProjectHandler)
	r.Get("/", controllers.GetProjectsHandler)
	r.Get("/{id}", controllers.GetProjectByIDHandler)
	r.Put("/{id}", controllers.UpdateProjectHandler)
	r.Delete("/{id}", controllers.DeleteProjectHandler)

	return r
}
