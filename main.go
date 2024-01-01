package main

import (
	"github.com/pithawatnuckong/go-clean/configuration"
	"github.com/pithawatnuckong/go-clean/environment"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	config, finder := environment.NewEnvironment()
	database := configuration.NewDatabase(config.Database)
	logs := configuration.NewZapLogging(config.Logging, finder)

	defer func(database *gorm.DB, logger *zap.Logger) {
		postgres, _ := database.DB()

		_ = postgres.Close()
		_ = logger.Sync()
	}(database, logs)
}
