package repository

import (
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/user-service/internal/domain"
)

type IRepository interface {
	IsExist(data, OPT string) bool
	Create(data domain.User) (domain.User, *serviceError.Error)
}
