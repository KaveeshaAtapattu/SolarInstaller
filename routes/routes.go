package routes

import (
	"net/http"
	"SolarInstaller/routes/todoRoutes"
	// "SolarInstaller/routes/userRoutes"

	"github.com/go-chi/chi"
)

func InitRoutes() http.Handler {
	r := chi.NewRouter()

	// Mount specific routes for each module
	r.Mount("/todos", todoRoutes.TodoRouter())
	// r.Mount("/users", userRoutes.UserRouter())

	return r
}
