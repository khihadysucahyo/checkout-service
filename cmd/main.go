package main

import (
	"log"

	_orderHttpDelivery "github.com/khihadysucahyo/checkout-service/order/delivery/http"
	_orderUcase "github.com/khihadysucahyo/checkout-service/order/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	orderUsecase := _orderUcase.NewOrderUsecase()
	_orderHttpDelivery.NewOrderHandler(e, orderUsecase)

	log.Fatal(e.Start(":2022"))
}
