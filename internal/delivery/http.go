package delivery

import (
	"github.com/Gym-Partner/api_common/router"
	"github.com/Gym-Partner/user-service/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterRoutes(router *gin.Engine, deps *router.Dependencies) {
	newController := controller.New(deps.Database)

	v1Auth := router.Group(viper.GetString("API_PREFIX"))
	{
		v1Auth.GET("/user/get_all", newController.GetAll)
	}

	v1NoAuth := router.Group(viper.GetString("API_PREFIX"))
	{
		v1NoAuth.POST("/user/create", newController.Create)
	}
}
