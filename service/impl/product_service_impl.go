package impl

import (
	"context"
	"github.com/pithawatnuckong/go-clean/entity"
	"github.com/pithawatnuckong/go-clean/exception"
	"github.com/pithawatnuckong/go-clean/model"
	"github.com/pithawatnuckong/go-clean/repository"
	"github.com/pithawatnuckong/go-clean/service"
	"strings"
	"time"
)

type productServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository *repository.ProductRepository) service.ProductService {
	return productServiceImpl{productRepository: *productRepository}
}

func (service productServiceImpl) CreateProduct(ctx context.Context, request model.ProductCreateOrUpdateModel) (response *model.ProductCreateOrUpdateModel) {
	request.Name = strings.TrimSpace(request.Name)
	if request.ID != 0 || request.Name == "" || request.Price < 0.0 || request.Quantity < 0 || request.OwnerID == 0 {
		panic(exception.ValidationError{
			Message: "Invalid product detail.",
		})
	}

	product := entity.Product{
		Name:      request.Name,
		Price:     request.Price,
		Quantity:  request.Quantity,
		OwnerID:   request.OwnerID,
		CreatedAt: time.Now(),
	}

	productId := service.productRepository.Create(ctx, product)

	return &model.ProductCreateOrUpdateModel{
		ID:       productId,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
		OwnerID:  request.OwnerID,
	}
}
