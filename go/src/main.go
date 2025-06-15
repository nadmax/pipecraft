package main

import (
	"log"
	"net/http"
	"github.com/nadmax/pipecraft/src/api"
	"github.com/nadmax/pipecraft/src/middlewares"
	"github.com/nadmax/pipecraft/src/routes"
)

func main() {
	api := api.NewAPI()

	api.Store.Create("John Doe", "john@example.com")
	api.Store.Create("Jane Smith", "jane@example.com")

	mux := routes.SetupRoutes(api)
	handler := middlewares.LoggingMiddleware(mux)
	port := ":8002"

	log.Printf("Server starting on port %s", port)
	log.Printf("API endpoints:")
	log.Printf("  GET    /users       - Get all users")
	log.Printf("  POST   /users       - Create a new user")
	log.Printf("  GET    /users/{id}  - Get user by ID")
	log.Printf("  PUT    /users/{id}  - Update user by ID")
	log.Printf("  DELETE /users/{id}  - Delete user by ID")

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

