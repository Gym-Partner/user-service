package repository

import (
	"github.com/Gym-Partner/api-common/serviceError"
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
	// Returns the created user and nil on success, or an empty user and a repository error on failure.
	Create(data domain.User) (domain.User, *serviceError.Error)

	// GetAll retrieves all users in the database.
	// Returns all retrieved users and nil on success, or an empty users and a repository error on failure.
	GetAll() (domain.Users, *serviceError.Error)

	// GetOneByID retrieve one user in the database using his ID.
	// Return retrieve user and nil on success, or an empty user and a repository error on failure.
	GetOneByID(uid string) (domain.User, *serviceError.Error)

	// GetOneByEmail retrieve one user in the database using his email.
	// Return retrieve user and nil on success, or an empty user and a repository error on failure.
	GetOneByEmail(email string) (domain.User, *serviceError.Error)
}
