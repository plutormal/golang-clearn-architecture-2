package routes

import (
	"golang-crud-api/handlers"
	"golang-crud-api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(middleware.CORSMiddleware)

	// Apply logging middleware
	r.Use(middleware.LoggingMiddleware)

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// User routes
	api.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Health check route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK", "message": "API is running"}`))
	}).Methods("GET")

	return r
}
