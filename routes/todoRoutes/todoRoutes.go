package todoRoutes

import (
	"SolarInstaller/controllers"
	"net/http"

	"github.com/go-chi/chi"
)

func TodoRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", controllers.FetchTodosHandler)
	r.Get("/{id}", controllers.GetTodoByIDHandler)

	r.Post("/", controllers.CreateTodoHandler)
	r.Put("/{id}", controllers.UpdateTodoHandler)
	r.Delete("/{id}", controllers.DeleteTodoHandler)

	return r
}
