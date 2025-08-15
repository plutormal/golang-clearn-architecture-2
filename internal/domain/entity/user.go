package entity

import (
	"errors"
	"time"
)

// User represents the core user entity
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser creates a new user with validation
func NewUser(id int, name, email string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	if age < 0 {
		return nil, errors.New("age must be non-negative")
	}

	now := time.Now()
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update updates user fields
func (u *User) Update(name, email string, age int) {
	if name != "" {
		u.Name = name
	}
	if email != "" {
		u.Email = email
	}
	if age > 0 {
		u.Age = age
	}
	u.UpdatedAt = time.Now()
}

// Validate validates user data
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Age < 0 {
		return errors.New("age must be non-negative")
	}
	return nil
}
