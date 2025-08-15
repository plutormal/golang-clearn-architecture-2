package services

import (
	"errors"
	"golang-crud-api/models"
	"sync"
	"time"
)

type UserService struct {
	users  []models.User
	nextID int
	mutex  sync.RWMutex
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		users:  make([]models.User, 0),
		nextID: 1,
	}
}

// GetAllUsers returns all users
func (s *UserService) GetAllUsers() []models.User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	// Return a copy of the slice
	usersCopy := make([]models.User, len(s.users))
	copy(usersCopy, s.users)
	return usersCopy
}

// GetUserByID returns a user by ID
func (s *UserService) GetUserByID(id int) (*models.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	
	return nil, errors.New("user not found")
}

// CreateUser creates a new user
func (s *UserService) CreateUser(name, email string, age int) *models.User {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	user := models.User{
		ID:        s.nextID,
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	s.users = append(s.users, user)
	s.nextID++
	
	return &user
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id int, name, email string, age int) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	for i, user := range s.users {
		if user.ID == id {
			// Update only non-empty fields
			if name != "" {
				s.users[i].Name = name
			}
			if email != "" {
				s.users[i].Email = email
			}
			if age > 0 {
				s.users[i].Age = age
			}
			s.users[i].UpdatedAt = time.Now()
			
			return &s.users[i], nil
		}
	}
	
	return nil, errors.New("user not found")
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	for i, user := range s.users {
		if user.ID == id {
			// Remove the user from the slice
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	
	return errors.New("user not found")
}
