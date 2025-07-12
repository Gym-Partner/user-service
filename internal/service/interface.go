package service

import (
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/gin-gonic/gin"
)

// IService defines the business logic contract for user-related operations.
type IService interface {

	// Create handles the creation of a new domain.User from request body in *gin.Context
	// It validates the input, check if the user already exist, and creates the user if valid.
	// Return the created user and nil on success, or an empty user and a service error on failure.
	Create(ctx *gin.Context) (domain.User, *serviceError.Error)
}
