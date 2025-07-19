package delivery

import (
	"github.com/Gym-Partner/api-common/middlewares"
	"github.com/Gym-Partner/api-common/router"
	"github.com/Gym-Partner/user-service/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterRoutes(router *gin.Engine, deps *router.Dependencies) {
	newController := controller.New(deps.Database, deps.Rabbit, deps.Catalog)

	v1Auth := router.Group(viper.GetString("API_PREFIX"), middlewares.Auth())
	{
		v1Auth.GET("/user/get/all", newController.GetAll)
		v1Auth.GET("/user/get/one/id", newController.GetOneByID)
		v1Auth.GET("/user/get/one/email", newController.GetOneByEmail)
	}

	v1NoAuth := router.Group(viper.GetString("API_PREFIX"))
	{
		v1NoAuth.POST("/user/create", newController.Create)
	}
}
