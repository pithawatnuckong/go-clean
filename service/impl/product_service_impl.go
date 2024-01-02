package impl

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pithawatnuckong/go-clean/configuration"
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
	redisClient       *redis.Client
}

func NewProductServiceImpl(productRepository *repository.ProductRepository, redisClient *redis.Client) service.ProductService {
	return productServiceImpl{productRepository: *productRepository, redisClient: redisClient}
}

func (service productServiceImpl) CreateProduct(ctx context.Context, request model.ProductCreateOrUpdateModel) *model.ProductCreateOrUpdateModel {
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

func (service productServiceImpl) FindProduct(ctx context.Context, id int) *model.ProductModel {
	if id <= 0 {
		panic(exception.ValidationError{
			Message: "Product ID must be greater than 0.",
		})
	}

	//product := service.productRepository.FindById(ctx, id)
	//if product == nil {
	//	panic(exception.ValidationError{
	//		Message: fmt.Sprintf("Produt ID %v not found.", id),
	//	})
	//}

	product := configuration.FindByIdAndSetCache[entity.Product](service.redisClient, ctx, "product", id, service.productRepository.FindById)

	return &model.ProductModel{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Quantity:  product.Quantity,
		OwnerID:   product.OwnerID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt.Time,
	}
}
