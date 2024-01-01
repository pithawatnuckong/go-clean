package main

import (
	"github.com/pithawatnuckong/go-clean/configuration"
	"github.com/pithawatnuckong/go-clean/environment"
	repository "github.com/pithawatnuckong/go-clean/repository/impl"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	config, finder := environment.NewEnvironment()
	database := configuration.NewDatabase(config.Database)
	logs := configuration.NewZapLogging(config.Logging, finder)

	// repositories
	productRepository := repository.NewProductRepositoryDBImpl(database)
	_ = productRepository

	//createEntity := entity.Product{
	//	Name:      "Zebra toy",
	//	Price:     50.00,
	//	Quantity:  2,
	//	OwnerID:   1,
	//	CreatedAt: time.Now(),
	//	UpdatedAt: sql.NullTime{Time: time.Now(), Valid: false},
	//	DeletedAt: sql.NullTime{Time: time.Now(), Valid: false},
	//}
	//productRepository.Create(context.Background(), createEntity)

	//foundProduct := productRepository.FindById(context.Background(), 4)
	//logs.Info("Find product by ID", zap.Any("product", foundProduct))

	//updateEntity := entity.Product{
	//	ID:        4,
	//	Name:      "Zebraa1",
	//	Price:     50.21,
	//	Quantity:  1,
	//	UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	//}
	//productRepository.Update(context.Background(), updateEntity)

	//productRepository.Delete(context.Background(), 4)

	//allProducts := productRepository.FindAll(context.Background())
	//logs.Info("Find all ", zap.Any("products", allProducts))

	defer func(database *gorm.DB, logger *zap.Logger) {
		postgres, _ := database.DB()

		_ = postgres.Close()
		_ = logger.Sync()
	}(database, logs)
}
