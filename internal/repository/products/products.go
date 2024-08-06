package products

import (
	"github.com/gatotbima1104/my-product/internal/models"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func GetRepo(db *gorm.DB) Repository {
	return &productRepo{
		db: db,
	}
}

func (pr *productRepo) GetProducts(productType string) ([]models.Products, error) {
	var products []models.Products

	if err := pr.db.Where(models.Products{Type: models.ProductsType(productType)}).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
