package main

import (
	"github.com/Gym-Partner/api-common/config"
	"github.com/Gym-Partner/api-common/router"
	"github.com/Gym-Partner/user-service/internal/delivery"
	"github.com/Gym-Partner/user-service/internal/domain"
)

func main() {
	// Initialize service config with options
	conf := config.InitConfig(config.Options{
		EnableDatabase: true,
		EnableRabbitMQ: true,
		Migrations:     []any{domain.MigrateUser{}},
		IsTest:         false,
	})

	// Initialize router with dependencies
	r := router.InitRouter(router.Options{
		Deps: &router.Dependencies{
			Database: conf.Database,
			Rabbit:   conf.RabbitMQ,
		},
		RegisterRoutes: delivery.RegisterRoutes,
	})

	// Run the service
	conf.Run(r)
}
