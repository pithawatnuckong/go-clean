package configuration

import "github.com/gofiber/fiber/v2/middleware/logger"

func NewFiberLoggerConfiguration() logger.Config {
	return logger.Config{
		Format:     "[${ip}]:${port} ${locals:requestID} ${latency} ${status} - ${method} ${path} \n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}
}
