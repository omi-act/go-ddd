package repositories

import (
	"go-ddd/internal/domain/entities"
	"go-ddd/internal/domain/value_objects"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// FindByID retrieves a user by ID
	FindByID(userID *value_objects.UserID) (*entities.User, error)

	// FindAll retrieves all users with optional pagination
	FindAll(limit, offset int) ([]*entities.User, error)
}
