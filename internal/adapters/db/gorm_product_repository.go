package db

import (
	"github.com/floxydio/boiler-fiber/internal/domain"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(dbClient *gorm.DB) domain.ProductRepository {
	return &GormProductRepository{
		db: dbClient,
	}
}

func (r *GormProductRepository) FindAll() ([]domain.Products, error) {
	var product []domain.Products
	err := r.db.Model(&product).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *GormProductRepository) FindById(id uint64) (domain.Products, error) {
	var productId domain.Products

	err := r.db.Model(&productId).Where("id = ?", id).First(&productId).Error

	if err != nil {
		return domain.Products{}, err
	}

	return productId, nil
}

func (r *GormProductRepository) Save(data *domain.ProductForm) error {
	err := r.db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}
