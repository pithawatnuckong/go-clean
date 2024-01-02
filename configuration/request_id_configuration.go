package configuration

import (
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

func NewRequestIdConfiguration() requestid.Config {
	return requestid.Config{
		Header:     "X-Request-ID",
		Generator:  utils.UUIDv4,
		ContextKey: "requestID",
	}
}
