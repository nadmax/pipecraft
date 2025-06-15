package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/nadmax/pipecraft/src/models"
	"github.com/nadmax/pipecraft/src/store"
)

type API struct {
	Store *store.UserStore
}

func NewAPI() *API {
	return &API{
		Store: store.NewUserStore(),
	}
}

func (api *API) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (api *API) respondError(w http.ResponseWriter, status int, message string) {
	api.respondJSON(w, status, map[string]string{"error": message})
}

func (api *API) extractUserID(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		return 0, http.ErrMissingFile
	}

	return strconv.Atoi(parts[1])
}

func (api *API) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users := api.Store.GetAll()
		api.respondJSON(w, http.StatusOK, users)
	case http.MethodPost:
		var req models.UserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			api.respondError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		if req.Name == "" || req.Email == "" {
			api.respondError(w, http.StatusBadRequest, "Name and email are required")
			return
		}
		user := api.Store.Create(req.Name, req.Email)
		api.respondJSON(w, http.StatusCreated, user)
	default:
		api.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (api *API) UserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := api.extractUserID(r.URL.Path)
	if err != nil {
		api.respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, exists := api.Store.GetByID(id)
		if !exists {
			api.respondError(w, http.StatusNotFound, "User not found")
			return
		}
		api.respondJSON(w, http.StatusOK, user)
	case http.MethodPut:
		var req models.UserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			api.respondError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}
		user, exists := api.Store.Update(id, req.Name, req.Email)
		if !exists {
			api.respondError(w, http.StatusNotFound, "User not found")
			return
		}
		api.respondJSON(w, http.StatusOK, user)
	case http.MethodDelete:
		if !api.Store.Delete(id) {
			api.respondError(w, http.StatusNotFound, "User not found")
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		api.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
