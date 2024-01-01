package configuration

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pithawatnuckong/go-clean/exception"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
