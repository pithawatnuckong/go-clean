package impl

import (
	"context"
	"github.com/pithawatnuckong/go-clean/entity"
	"github.com/pithawatnuckong/go-clean/exception"
	"github.com/pithawatnuckong/go-clean/repository"
	"gorm.io/gorm"
	"time"
)

type productRepositoryDBImpl struct {
	database *gorm.DB
}

func NewProductRepositoryDBImpl(database *gorm.DB) repository.ProductRepository {
	exception.PanicLogging(database.AutoMigrate(&entity.Product{}))
	return productRepositoryDBImpl{database: database}
}

func (repository productRepositoryDBImpl) Create(ctx context.Context, product entity.Product) (id int) {
	err := repository.database.WithContext(ctx).Create(&product).Error
	exception.PanicLogging(err)
	return product.ID
}

func (repository productRepositoryDBImpl) Update(ctx context.Context, product entity.Product) {
	err := repository.database.WithContext(ctx).Where("product_id=?", product.ID).Updates(&product).Error
	exception.PanicLogging(err)
}

func (repository productRepositoryDBImpl) Delete(ctx context.Context, id int) {
	err := repository.database.WithContext(ctx).Where("product_id=?", id).Model(&entity.Product{}).Update("deleted_at", time.Now()).Error
	exception.PanicLogging(err)
}

func (repository productRepositoryDBImpl) FindAll(ctx context.Context) (products []entity.Product) {
	err := repository.database.WithContext(ctx).Where("deleted_at IS null").Find(&products).Error
	exception.PanicLogging(err)
	return products
}

func (repository productRepositoryDBImpl) FindById(ctx context.Context, id int) (product *entity.Product) {
	err := repository.database.WithContext(ctx).Where("product_id=?", id).Where("deleted_at IS null").First(&product).Error
	exception.PanicLogging(err)
	return product
}
