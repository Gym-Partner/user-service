package controller

import (
	"github.com/Gym-Partner/api_common/database"
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/Gym-Partner/api_common/utils"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/Gym-Partner/user-service/internal/repository"
	"github.com/Gym-Partner/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	IService service.IService
}

func New(db *database.Database) *Controller {
	repo := repository.New(db.Handler, db.Logger)
	svc := service.New(repo, utils.NewUtils[domain.User]())

	return &Controller{IService: svc}
}

func (c *Controller) Create(ctx *gin.Context) {
	user, err := c.IService.Create(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}

	ctx.JSON(serviceError.HttpCode201.ToInt(), user.Response())
}
