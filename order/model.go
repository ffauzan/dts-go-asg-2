package order

import (
	"context"
	"time"
)

type Order struct {
	ID           int
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}

type Item struct {
	ID          int
	Code        string
	Description string
	Quantity    int
	OrderID     int
}

type OrderService interface {
	CreateOrder(ctx context.Context, order Order) (Order, error)
	GetOrderByID(ctx context.Context, orderID int) (Order, error)
	GetOrders(ctx context.Context) ([]Order, error)
	UpdateOrder(ctx context.Context, order Order) (Order, error)
	DeleteOrder(ctx context.Context, orderID int) error
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order Order) (Order, error)
	GetOrderByID(ctx context.Context, orderID int) (Order, error)
	GetOrders(ctx context.Context) ([]Order, error)
	UpdateOrder(ctx context.Context, order Order) (Order, error)
	DeleteOrder(ctx context.Context, orderID int) error
}
