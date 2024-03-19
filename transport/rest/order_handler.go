package rest

import (
	"asg-2/order"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CreateOrderRequest struct {
	CustomerName string                   `json:"customerName" validate:"required"`
	Items        []ItemCreateOrderRequest `json:"items" validate:"required,dive"`
}

type ItemCreateOrderRequest struct {
	ItemCode    string `json:"itemCode" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
}

type orderHandler struct {
	service order.OrderService
}

func NewOrderHandler(service order.OrderService) *orderHandler {
	return &orderHandler{
		service: service,
	}
}

func (h *orderHandler) CreateOrder(c echo.Context) error {
	req := new(CreateOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	o := order.Order{
		CustomerName: req.CustomerName,
		Items:        []order.Item{},
	}
	for _, item := range req.Items {
		o.Items = append(o.Items, order.Item{
			Code:        item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	createdOrder, err := h.service.CreateOrder(c.Request().Context(), o)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, BaseResponse{
		Status:  http.StatusCreated,
		Message: "Order created",
		Data:    createdOrder,
	})
}

func (h *orderHandler) GetOrders(c echo.Context) error {
	orders, err := h.service.GetOrders(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  http.StatusOK,
		Message: "Orders",
		Data:    orders,
	})
}

func (h *orderHandler) DeleteOrder(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid order ID",
			Data:    nil,
		})
	}

	err = h.service.DeleteOrder(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  http.StatusOK,
		Message: "Order deleted",
		Data:    nil,
	})
}

func (h *orderHandler) UpdateOrder(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid order ID",
			Data:    nil,
		})
	}

	req := new(CreateOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	o := order.Order{
		ID:           id,
		CustomerName: req.CustomerName,
		Items:        []order.Item{},
	}
	for _, item := range req.Items {
		o.Items = append(o.Items, order.Item{
			Code:        item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	updatedOrder, err := h.service.UpdateOrder(c.Request().Context(), o)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Status:  http.StatusOK,
		Message: "Order updated",
		Data:    updatedOrder,
	})
}
