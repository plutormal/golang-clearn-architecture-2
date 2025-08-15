package repository

import "golang-crud-api/internal/domain/entity"

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetAll() ([]*entity.User, error)
	GetByID(id int) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id int) error
	GetNextID() int
}
