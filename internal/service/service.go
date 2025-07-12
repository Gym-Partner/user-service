package service

import (
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/api_common/utils"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/Gym-Partner/user-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type Service struct {
	IRepository repository.IRepository
	Utils       *utils.Utils[domain.User]
}

func (s Service) Create(ctx *gin.Context) (domain.User, *serviceError.Error) {
	//TODO implement me
	panic("implement me")
}
