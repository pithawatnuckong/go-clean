package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/pithawatnuckong/go-clean/configuration"
	"github.com/pithawatnuckong/go-clean/controller"
	"github.com/pithawatnuckong/go-clean/environment"
	"github.com/pithawatnuckong/go-clean/exception"
	"github.com/pithawatnuckong/go-clean/logs"
	repository "github.com/pithawatnuckong/go-clean/repository/impl"
	service "github.com/pithawatnuckong/go-clean/service/impl"
	"gorm.io/gorm"
)

// TODO implement error log

func main() {
	config, finder := environment.NewEnvironment()
	database := configuration.NewDatabase(config.Database)
	redisClient := configuration.NewRedis(config.Redis)

	defer terminate(database, redisClient)

	logs.NewCustomerLogger(config.Logging, finder)

	// repositories
	productRepository := repository.NewProductRepositoryDBImpl(database)

	// services
	productService := service.NewProductServiceImpl(&productRepository)

	// controllers
	productController := controller.NewProductController(&productService)

	// set-up fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New(configuration.NewCorsConfiguration()))
	app.Use(requestid.New(configuration.NewRequestIdConfiguration()))
	app.Use(logger.New(configuration.NewFiberLoggerConfiguration()))

	// routes
	productController.Route(app)

	err := app.Listen(fmt.Sprintf(":%v", finder.Get("server.port")))
	exception.PanicLogging(err)
}

func terminate(database *gorm.DB, redisClient *redis.Client) {
	postgres, _ := database.DB()

	_ = postgres.Close()
	_ = redisClient.Close()
}
