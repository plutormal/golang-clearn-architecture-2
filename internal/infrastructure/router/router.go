package router

import (
	"net/http"

	"golang-crud-api/internal/interface/controller"
	"golang-crud-api/middleware"

	"github.com/gorilla/mux"
)

// Router sets up all the routes for the application
func SetupRoutes(userController *controller.UserController) *mux.Router {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(middleware.CORSMiddleware)

	// Apply logging middleware
	r.Use(middleware.LoggingMiddleware)

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// User routes
	api.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	api.HandleFunc("/users", userController.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Health check route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK", "message": "API is running"}`))
	}).Methods("GET")

	return r
}
