package presenter

import (
	"encoding/json"
	"net/http"
)

// UserPresenter handles response formatting for user operations
type UserPresenter struct{}

// NewUserPresenter creates a new user presenter
func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

// Response represents the standard API response format
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PresentSuccess presents a successful response
func (p *UserPresenter) PresentSuccess(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := Response{
		Success: true,
		Data:    data,
		Message: message,
	}
	
	json.NewEncoder(w).Encode(response)
}

// PresentError presents an error response
func (p *UserPresenter) PresentError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := Response{
		Success: false,
		Error:   err.Error(),
	}
	
	json.NewEncoder(w).Encode(response)
}
