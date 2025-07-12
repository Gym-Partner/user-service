package repository

import (
	"fmt"
	"github.com/Gym-Partner/api_common/logger"
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/user-service/internal/constants"
	"github.com/Gym-Partner/user-service/internal/domain"
	"gorm.io/gorm"
	"strings"
)

// Repository provides access to the user data stored in the database.
type Repository struct {
	DB  *gorm.DB
	Log *logger.Logger
}

// New creates and returns a new Repository instance
// using the provided GORM database handler and logger
func New(db *gorm.DB, logg *logger.Logger) *Repository {
	return &Repository{
		DB:  db,
		Log: logg,
	}
}

// IsExist implements IRepository.IsExist
func (r *Repository) IsExist(data, OPT string) bool {
	var user domain.User
	var queryColumn string

	switch strings.ToLower(OPT) {
	case "ID":
		queryColumn = "id"
	case "EMAIL":
		queryColumn = "email"
	}

	if raw := r.DB.
		Where(queryColumn+" = ?", data).
		Find(&user); raw.Error != nil {
		r.Log.Error(raw.Error.Error())
		return false
	}

	if user.ID == "" {
		r.Log.Error(constants.ServiceErrDBUserNotFound)
		return false
	} else {
		return true
	}
}

// Create implements IRepository.Create
func (r *Repository) Create(data domain.User) (domain.User, *serviceError.Error) {
	if raw := r.DB.
		Create(&data); raw.Error != nil {
		r.Log.Error(raw.Error.Error())
		return domain.User{}, serviceError.New(
			serviceError.HttpCode500,
			fmt.Sprintf(constants.ServiceErrAppDBCreateUser, data.Email),
			serviceError.WithOriginal(raw.Error))
	}
	return data, nil
}
