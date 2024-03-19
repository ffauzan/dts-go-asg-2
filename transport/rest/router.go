package rest

import (
	"asg-2/order"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewRouter(orderService order.OrderService) http.Handler {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency}\n",
	}))

	e.Validator = NewValidator()

	// Order routes
	orderHandler := NewOrderHandler(orderService)
	e.POST("/orders", orderHandler.CreateOrder)
	e.GET("/orders", orderHandler.GetOrders)
	e.DELETE("/orders/:id", orderHandler.DeleteOrder)
	e.PUT("/orders/:id", orderHandler.UpdateOrder)

	return e
}
