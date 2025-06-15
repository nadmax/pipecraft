package routes

import (
	"net/http"
	"github.com/nadmax/pipecraft/src/api"
)

func SetupRoutes(api *api.API) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", api.UsersHandler)
	mux.HandleFunc("/users/", api.UserHandler)

	return mux
}
