package main

import (
	"log"
	"net/http"

	"golang-crud-api/internal/infrastructure/repository"
	"golang-crud-api/internal/infrastructure/router"
	"golang-crud-api/internal/interface/controller"
	"golang-crud-api/internal/interface/presenter"
	"golang-crud-api/internal/usecase"
)

func main() {
	// Initialize dependencies following Clean Architecture layers
	
	// Infrastructure layer - Repository
	userRepo := repository.NewMemoryUserRepository()
	
	// Use case layer - Business logic
	userUseCase := usecase.NewUserUseCase(userRepo)
	
	// Interface adapters layer - Presenter and Controller
	userPresenter := presenter.NewUserPresenter()
	userController := controller.NewUserController(userUseCase, userPresenter)
	
	// Infrastructure layer - Router
	appRouter := router.SetupRoutes(userController)

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Printf("Clean Architecture API endpoints:")
	log.Printf("  GET    /health              - Health check")
	log.Printf("  GET    /api/v1/users        - Get all users")
	log.Printf("  GET    /api/v1/users/{id}   - Get user by ID")
	log.Printf("  POST   /api/v1/users        - Create new user")
	log.Printf("  PUT    /api/v1/users/{id}   - Update user")
	log.Printf("  DELETE /api/v1/users/{id}   - Delete user")

	if err := http.ListenAndServe(port, appRouter); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
