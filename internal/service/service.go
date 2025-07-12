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

// Service handles user-related business logic
type Service struct {
	IRepository repository.IRepository
	Utils       *utils.Utils[domain.User]
}

// New creates and return a new Service instance with its dependencies
func New(repo repository.IRepository, tools *utils.Utils[domain.User]) *Service {
	return &Service{
		IRepository: repo,
		Utils:       tools,
	}
}

// Create implements IService.Create
func (s *Service) Create(ctx *gin.Context) (domain.User, *serviceError.Error) {
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

	userPtr.GenerateId()
	user, err := s.IRepository.Create(*userPtr)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
