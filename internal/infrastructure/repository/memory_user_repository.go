package repository

import (
	"errors"
	"sync"

	"golang-crud-api/internal/domain/entity"
	"golang-crud-api/internal/domain/repository"
)

// MemoryUserRepository implements UserRepository using in-memory storage
type MemoryUserRepository struct {
	users  map[int]*entity.User
	nextID int
	mutex  sync.RWMutex
}

// NewMemoryUserRepository creates a new memory user repository
func NewMemoryUserRepository() repository.UserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]*entity.User),
		nextID: 1,
	}
}

// GetAll returns all users
func (r *MemoryUserRepository) GetAll() ([]*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	users := make([]*entity.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	
	return users, nil
}

// GetByID returns a user by ID
func (r *MemoryUserRepository) GetByID(id int) (*entity.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	
	return user, nil
}

// Create creates a new user
func (r *MemoryUserRepository) Create(user *entity.User) (*entity.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	r.users[user.ID] = user
	if user.ID >= r.nextID {
		r.nextID = user.ID + 1
	}
	
	return user, nil
}

// Update updates an existing user
func (r *MemoryUserRepository) Update(user *entity.User) (*entity.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[user.ID]; !exists {
		return nil, errors.New("user not found")
	}
	
	r.users[user.ID] = user
	return user, nil
}

// Delete deletes a user by ID
func (r *MemoryUserRepository) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	
	delete(r.users, id)
	return nil
}

// GetNextID returns the next available ID
func (r *MemoryUserRepository) GetNextID() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.nextID
}
