package service

import (
	"github.com/Gym-Partner/api-common/errs"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/gin-gonic/gin"
)

// IService defines the business logic contract for user-related operations.
type IService interface {

	// Create handles the creation of a new domain.User from request body in *gin.Context
	// It validates the input, check if the user already exist, and creates the user if valid.
	// Return the created user and nil on success, or an empty user and a service error on failure.
	Create(ctx *gin.Context) (domain.User, *errs.Error)

	// GetAll manages the retrieval of all users from the database.
	// Returns all users and nil on success, or an empty users and a service error on failure.
	GetAll() (users domain.Users, err *errs.Error)

	// GetOneByID manages the retrieval of one user from the database with his ID in user's token.
	// Return the user and nil on success, or an empty user and a service error on failure.
	GetOneByID(ctx *gin.Context) (user domain.User, err *errs.Error)

	// GetOneByEmail manages the retrieval of on user from the database with his email.
	// Return the user and nil on success, or an empty user and a service error on failure.
	GetOneByEmail(ctx *gin.Context) (domain.User, *errs.Error)
}
