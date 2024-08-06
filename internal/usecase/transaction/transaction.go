package transaction

import (
	"github.com/gatotbima1104/my-product/internal/models"
	"github.com/gatotbima1104/my-product/internal/repository/products"
)

type transUsecase struct {
	productRepo products.Repository
}

func GetUsecase(productRepo products.Repository) Usecase {
	return &transUsecase{
		productRepo: productRepo,
	}
}

func (t *transUsecase) GetProducts(productType string) ([]models.Products, error) {
	return t.productRepo.GetProducts(productType)
}
