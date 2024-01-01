package repository

import (
	"context"
	"github.com/pithawatnuckong/go-clean/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, product entity.Product) (id int)
	Update(ctx context.Context, product entity.Product)
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) (products []entity.Product)
	FindById(ctx context.Context, id int) (product *entity.Product)
}
