package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetProducts(c echo.Context) error {
	productType := c.FormValue("product_type")

	products, err := h.transUsecase.GetProducts(productType)
	if err != nil {

		fmt.Printf("Error in %s", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": products,
	})
}
