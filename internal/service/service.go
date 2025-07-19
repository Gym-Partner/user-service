package service

import (
	"github.com/Gym-Partner/api-common/errs"
	"github.com/Gym-Partner/api-common/rabbitmq"
	"github.com/Gym-Partner/api-common/status"
	"github.com/Gym-Partner/api-common/utils"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/Gym-Partner/user-service/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Service handles user-related business logic
type Service struct {
	IRepository repository.IRepository
	Rabbit      *rabbitmq.RabbitMQ
	Catalog     *status.StatusCatalog
	Utils       *utils.Utils[domain.User]
}

// New creates and return a new Service instance with its dependencies
func New(repo repository.IRepository, rabbit *rabbitmq.RabbitMQ, catalog *status.StatusCatalog, util *utils.Utils[domain.User]) *Service {
	return &Service{
		IRepository: repo,
		Rabbit:      rabbit,
		Catalog:     catalog,
		Utils:       util,
	}
}

// Create implements IService.Create
func (s *Service) Create(ctx *gin.Context) (domain.User, *errs.Error) {
	userPtr, err := s.Utils.InjectBodyInModel(ctx)
	if err != nil {
		return domain.User{}, err
	}

	exist := s.IRepository.IsExist(userPtr.Email, "email")
	if exist {
		return domain.User{}, errs.New(s.Catalog, "USER_ALREADY_EXIST", nil, userPtr.Email)
	}

	userPtr.GenerateId()
	if err = userPtr.HashPassword(s.Utils.HashPassword); err != nil {
		return domain.User{}, err
	}

	user, err := s.IRepository.Create(*userPtr)
	if err != nil {
		return domain.User{}, err
	}

	// RabbitMQ Part
	if err = s.Rabbit.PublishEvent(rabbitmq.EventUserCreated, viper.GetString("SERVICE_NAME"), domain.User{
		ID:    user.ID,
		Email: user.Email,
	}); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// GetAll implements IService.GetAll
func (s *Service) GetAll() (users domain.Users, err *errs.Error) {
	users, err = s.IRepository.GetAll()
	return
}

// GetOneByID implements IService.GetOneByID
func (s *Service) GetOneByID(ctx *gin.Context) (user domain.User, err *errs.Error) {
	uid, _ := ctx.Get("uid")

	user, err = s.IRepository.GetOneByID(uid.(string))
	return
}

// GetOneByEmail implements IService.GetOneByEmail
func (s *Service) GetOneByEmail(ctx *gin.Context) (domain.User, *errs.Error) {
	data, err := s.Utils.InjectBodyInModel(ctx)
	if err != nil {
		return domain.User{}, err
	}

	user, err := s.IRepository.GetOneByEmail(data.Email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
