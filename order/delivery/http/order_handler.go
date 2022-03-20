package http

import (
	"net/http"

	"github.com/khihadysucahyo/checkout-service/domain"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderUseCase domain.OrderUsecase
}

func NewOrderHandler(e *echo.Echo, cartUseCase domain.OrderUsecase) {
	handler := &OrderHandler{
		OrderUseCase: cartUseCase,
	}

	e.POST("/order", handler.DoOrder)
}

func (h *OrderHandler) DoOrder(c echo.Context) error {

	requestItems := []domain.RequestItem{}

	if err := c.Bind(&requestItems); err != nil {
		return err
	}

	order, err := h.OrderUseCase.DoOrder(requestItems)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"errors": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, order)
}
