package transaction

import "github.com/gatotbima1104/my-product/internal/models"

type Usecase interface {
	GetProducts(productType string) ([]models.Products, error)
}
