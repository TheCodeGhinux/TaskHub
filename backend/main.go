package main

import (
	"github.com/TheCodeGhinux/TaskHub/taskhub/internal/models/migrations"
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/config"
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/repository/db/postgres"
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/routing"
)

// @title Naija prime streaming service microservice api
// @version 1.0
// @description Naija prime streaming service microservice api built with Gin
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /api/v1
// @schemes http

func main() {
	// Initialize the logger
	// logger := utils.InitLogger()

	configs := config.LoadConfig()
	_ = configs

	postgres.ConnectDb(configs.DB)
	migrations.MigrateDb()
	routing.Route()
}
