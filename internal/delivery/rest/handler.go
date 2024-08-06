package rest

import "github.com/gatotbima1104/my-product/internal/usecase/transaction"

type handler struct{
	transUsecase transaction.Usecase
}

func NewHandler(transUsecase transaction.Usecase) *handler {
	return &handler{
		transUsecase: transUsecase,
	}
}