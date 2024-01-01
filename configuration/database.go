package configuration

import (
	"fmt"
	"github.com/pithawatnuckong/go-clean/environment"
	"github.com/pithawatnuckong/go-clean/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewDatabase(environment environment.DatabaseEnv) *gorm.DB {
	datasource := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		environment.Host,
		environment.Port,
		environment.Username,
		environment.Password,
		environment.Name,
		environment.SslMode,
	)

	loggerDB := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dial := postgres.Open(datasource)

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: loggerDB,
	})
	exception.PanicLogging(err)

	return db
}
