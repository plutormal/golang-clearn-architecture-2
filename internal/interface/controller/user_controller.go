package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-crud-api/internal/interface/presenter"
	"golang-crud-api/internal/usecase"

	"github.com/gorilla/mux"
)

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   int    `json:"age,omitempty"`
}

// UserController handles HTTP requests for user operations
type UserController struct {
	userUseCase *usecase.UserUseCase
	presenter   *presenter.UserPresenter
}

// NewUserController creates a new user controller
func NewUserController(userUseCase *usecase.UserUseCase, presenter *presenter.UserPresenter) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		presenter:   presenter,
	}
}

// GetAllUsers handles GET /users
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userUseCase.GetAllUsers()
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusInternalServerError)
		return
	}

	c.presenter.PresentSuccess(w, users, "", http.StatusOK)
}

// GetUser handles GET /users/{id}
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusBadRequest)
		return
	}
	
	user, err := c.userUseCase.GetUserByID(id)
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusNotFound)
		return
	}
	
	c.presenter.PresentSuccess(w, user, "", http.StatusOK)
}

// CreateUser handles POST /users
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.presenter.PresentError(w, err, http.StatusBadRequest)
		return
	}
	
	user, err := c.userUseCase.CreateUser(req.Name, req.Email, req.Age)
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusBadRequest)
		return
	}
	
	c.presenter.PresentSuccess(w, user, "User created successfully", http.StatusCreated)
}

// UpdateUser handles PUT /users/{id}
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusBadRequest)
		return
	}
	
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.presenter.PresentError(w, err, http.StatusBadRequest)
		return
	}
	
	user, err := c.userUseCase.UpdateUser(id, req.Name, req.Email, req.Age)
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusNotFound)
		return
	}
	
	c.presenter.PresentSuccess(w, user, "User updated successfully", http.StatusOK)
}

// DeleteUser handles DELETE /users/{id}
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusBadRequest)
		return
	}
	
	err = c.userUseCase.DeleteUser(id)
	if err != nil {
		c.presenter.PresentError(w, err, http.StatusNotFound)
		return
	}
	
	c.presenter.PresentSuccess(w, nil, "User deleted successfully", http.StatusOK)
}
