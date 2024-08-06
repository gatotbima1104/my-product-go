package main

import (
	"github.com/gatotbima1104/my-product/internal/config"
	"github.com/gatotbima1104/my-product/internal/delivery/rest"
	"github.com/gatotbima1104/my-product/internal/repository/products"
	"github.com/gatotbima1104/my-product/internal/usecase/transaction"

	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()

	// db
	db := config.ConnectDatabase()

	// Repo > usecase > handler
	productRepo := products.GetRepo(db)
	transUsecase := transaction.GetUsecase(productRepo)
	handler := rest.NewHandler(transUsecase)

	rest.LoadRoutes(e, handler)

	e.Logger.Fatal(e.Start(":8080"))
}
