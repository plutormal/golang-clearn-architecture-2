package usecase

import (
	"golang-crud-api/internal/domain/entity"
	"golang-crud-api/internal/domain/repository"
)

// UserUseCase contains user-related business logic
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// GetAllUsers retrieves all users
func (uc *UserUseCase) GetAllUsers() ([]*entity.User, error) {
	return uc.userRepo.GetAll()
}

// GetUserByID retrieves a user by ID
func (uc *UserUseCase) GetUserByID(id int) (*entity.User, error) {
	return uc.userRepo.GetByID(id)
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(name, email string, age int) (*entity.User, error) {
	// Business logic: validate input and create user
	id := uc.userRepo.GetNextID()
	user, err := entity.NewUser(id, name, email, age)
	if err != nil {
		return nil, err
	}

	return uc.userRepo.Create(user)
}

// UpdateUser updates an existing user
func (uc *UserUseCase) UpdateUser(id int, name, email string, age int) (*entity.User, error) {
	// First, get the existing user
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update the user
	user.Update(name, email, age)

	// Validate after update
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return uc.userRepo.Update(user)
}

// DeleteUser deletes a user by ID
func (uc *UserUseCase) DeleteUser(id int) error {
	// Business logic: check if user exists before deleting
	_, err := uc.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return uc.userRepo.Delete(id)
}
