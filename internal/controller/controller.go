package controller

import (
	"github.com/Gym-Partner/api-common/database"
	"github.com/Gym-Partner/api-common/rabbitmq"
	"github.com/Gym-Partner/api-common/status"
	"github.com/Gym-Partner/api-common/utils"
	"github.com/Gym-Partner/user-service/internal/domain"
	"github.com/Gym-Partner/user-service/internal/repository"
	"github.com/Gym-Partner/user-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller provides HTTP handlers for user-related operations.
// It delegates business logic to the injected IService implementation.
type Controller struct {
	IService service.IService
}

// New creates a new instance of Controller by writing up to repository, service,
// utility and rabbitMQ components using the given database connection.
func New(db *database.Database, rabbit *rabbitmq.RabbitMQ, catalog *status.StatusCatalog) *Controller {
	repo := repository.New(db.Handler, catalog)
	svc := service.New(repo, rabbit, catalog, utils.NewUtils[domain.User]())

	return &Controller{IService: svc}
}

func (c *Controller) Create(ctx *gin.Context) {
	user, err := c.IService.Create(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}

	ctx.JSON(http.StatusCreated, user.Response())
}

func (c *Controller) GetAll(ctx *gin.Context) {
	users, err := c.IService.GetAll()
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *Controller) GetOneByID(ctx *gin.Context) {
	user, err := c.IService.GetOneByID(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *Controller) GetOneByEmail(ctx *gin.Context) {
	user, err := c.IService.GetOneByEmail(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
