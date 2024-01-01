package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pithawatnuckong/go-clean/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validateError := err.(ValidationError)
	if validateError {
		message := err.Error()
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponseModel{
			Code:    "P400",
			Message: "Bad Request",
			Data:    message,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponseModel{
		Code:    "P500",
		Message: "General Error",
		Data:    err.Error(),
	})
}
