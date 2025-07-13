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

// Controller provides HTTP handlers for user-related operations.
// It delegates business logic to the injected IService implementation.
type Controller struct {
	IService service.IService
}

// New creates a new instance of Controller by writing up to repository, service,
// utility and rabbitMQ components using the given database connection.
func New(db *database.Database) *Controller {
	repo := repository.New(db.Handler, db.Logger)
	svc := service.New(repo, utils.NewUtils[domain.User]())

	return &Controller{IService: svc}
}

// Create handles the HTTP POST request to create a new user.
// It delegates the creation to the service layer and returns the result
// as a JSON response with the appropriate HTTP status code.
func (c *Controller) Create(ctx *gin.Context) {
	user, err := c.IService.Create(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}

	ctx.JSON(serviceError.HttpCode201.ToInt(), user.Response())
}

// GetAll handles the HTTP GET request to retrieve all users.
// It calls the services layer to fetch the user list and returns it
// as a JSON response with the appropriate HTTP status code.
func (c *Controller) GetAll(ctx *gin.Context) {
	users, err := c.IService.GetAll()
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}
	ctx.JSON(serviceError.HttpCode200.ToInt(), users)
}

// GetOne handles the HTTP GET request to retrieve user by his email in token.
// It calls the services layer to fetch the user and returns it
// as a JSON response with the appropriate HTTP status code.
func (c *Controller) GetOne(ctx *gin.Context) {
	user, err := c.IService.GetOne(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Response())
		return
	}

	ctx.JSON(serviceError.HttpCode200.ToInt(), user)
}
