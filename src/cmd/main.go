package main

import (
	"github.com/tgmtime/golang-api-architecture/api"
	"github.com/tgmtime/golang-api-architecture/config"
	"github.com/tgmtime/golang-api-architecture/infra/cache"
	database "github.com/tgmtime/golang-api-architecture/infra/persistence/database"
	"github.com/tgmtime/golang-api-architecture/infra/persistence/migration"
	"github.com/tgmtime/golang-api-architecture/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up1()

	api.InitServer(cfg)
}
