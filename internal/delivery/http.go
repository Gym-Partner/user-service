package delivery

import (
	"github.com/Gym-Partner/api_common/router"
	"github.com/Gym-Partner/user-service/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterRoutes(router *gin.Engine, deps *router.Dependencies) {
	newController := controller.New(deps.Database)

	v1 := router.Group(viper.GetString("API_PREFIX"))
	{
		v1.POST("/user/create", newController.Create)
	}
}
