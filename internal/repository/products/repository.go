package products

import "github.com/gatotbima1104/my-product/internal/models"

type Repository interface {
	GetProducts(productType string) ([]models.Products, error)
}
