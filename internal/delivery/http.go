package delivery

import (
	"github.com/Gym-Partner/api_common/database"
	"github.com/Gym-Partner/user-service/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterRoutes(router *gin.Engine, db *database.Database) {
	newController := controller.New(db)

	v1 := router.Group(viper.GetString("API_PREFIX"))
	{
		v1.POST("/user/create", newController.Create)
	}
}
