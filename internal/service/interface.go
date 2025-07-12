package service

import (
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/gin-gonic/gin"
)

type IService interface {
	Create(ctx *gin.Context) (domain.User, *serviceError.Error)
}
