package configuration

import "github.com/gofiber/fiber/v2/middleware/cors"

func NewCorsConfiguration() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}
}
