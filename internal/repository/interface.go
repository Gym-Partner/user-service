package repository

import (
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/user-service/internal/domain"
)

// IRepository defines the contract for accessing user data
type IRepository interface {

	// IsExist check if a user exists in the database based on the provided value
	// and the given field option ("ID" or "EMAIL").
	// Return true if the user exist, false otherwise.
	// If an error occurs, it logs the error and returns false.
	IsExist(data, OPT string) bool

	// Create inserts a new user record into the database.
	// Returns the created user and nil on success, or an empty user and a service error on failure.
	Create(data domain.User) (domain.User, *serviceError.Error)
}
