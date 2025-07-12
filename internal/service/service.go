package service

import (
	"fmt"
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/api_common/utils"
	"github.com/Gym-Partner/user-service/internal/constants"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/Gym-Partner/user-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type Service struct {
	IRepository repository.IRepository
	Utils       *utils.Utils[domain.User]
}

func (s Service) Create(ctx *gin.Context) (domain.User, *serviceError.Error) {
	userPtr, err := s.Utils.InjectBodyInModel(ctx)
	if err != nil {
		return domain.User{}, err
	}

	exist := s.IRepository.IsExist(userPtr.Email, "email")
	if exist {
		return domain.User{}, serviceError.New(
			serviceError.HttpCode400,
			fmt.Sprintf(constants.ServiceErrAppINTUserAlreadyExist, userPtr.Email))
	}

	user, err := s.IRepository.Create(*userPtr)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
