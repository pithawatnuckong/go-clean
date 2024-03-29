package service

import (
	"context"
	"github.com/pithawatnuckong/go-clean/model"
)

type ProductService interface {
	CreateProduct(ctx context.Context, request model.ProductCreateOrUpdateModel) (response *model.ProductCreateOrUpdateModel)
	FindProduct(ctx context.Context, id int) (response *model.ProductModel)
}
