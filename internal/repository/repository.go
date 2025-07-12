package repository

import (
	"github.com/Gym-Partner/api_common/logger"
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/user-service/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	DB  *gorm.DB
	Log *logger.Logger
}

func (r Repository) IsExist(data, OPT string) bool {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Create(data domain.User) (domain.User, *serviceError.Error) {
	//TODO implement me
	panic("implement me")
}
