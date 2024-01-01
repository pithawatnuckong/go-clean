package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pithawatnuckong/go-clean/exception"
	"github.com/pithawatnuckong/go-clean/model"
	"github.com/pithawatnuckong/go-clean/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{productService: *productService}
}

func (controller ProductController) Route(app *fiber.App) {
	router := app.Group("/api/v1/products")
	router.Post("/create", controller.CreateProduct)
}

func (controller ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var body model.ProductCreateOrUpdateModel
	err := ctx.BodyParser(&body)
	exception.PanicLogging(err)

	response := controller.productService.CreateProduct(ctx.Context(), body)

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponseModel{
		Code:    "P201",
		Message: "Created",
		Data:    response,
	})
}
