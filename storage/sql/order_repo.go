package sql

import (
	"asg-2/order"
	"asg-2/storage/sql/db"
	"context"

	"github.com/jackc/pgx/v5"
)

type orderRepo struct {
	store   *store
	Conn    *pgx.Conn
	Queries *db.Queries
}

func NewOrderRepo(s *store) order.OrderRepo {
	return &orderRepo{
		store:   s,
		Conn:    s.Conn,
		Queries: s.Queries,
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, o order.Order) (order.Order, error) {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return order.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := r.Queries.WithTx(tx)
	dbOrder, err := qtx.CreateOrder(ctx, o.CustomerName)
	if err != nil {
		return order.Order{}, err
	}

	for _, item := range o.Items {
		item.OrderID = int(dbOrder.ID)
		_, err := qtx.CreateItem(ctx, db.CreateItemParams{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    int32(item.Quantity),
			OrderID:     int32(item.OrderID),
		})
		if err != nil {
			return order.Order{}, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return order.Order{}, err
	}

	return r.GetOrderByID(ctx, int(dbOrder.ID))
}

func (r *orderRepo) GetOrderByID(ctx context.Context, orderID int) (order.Order, error) {
	dbOrder, err := r.Queries.GetOrder(ctx, int32(orderID))
	if err != nil {
		return order.Order{}, err
	}

	dbItems, err := r.Queries.GetItemByOrderID(ctx, int32(orderID))
	if err != nil {
		return order.Order{}, err
	}

	items := []order.Item{}
	for _, dbItem := range dbItems {
		items = append(items, order.Item{
			ID:          int(dbItem.ID),
			Code:        dbItem.Code,
			Description: dbItem.Description,
			Quantity:    int(dbItem.Quantity),
			OrderID:     int(dbItem.OrderID),
		})
	}

	return order.Order{
		ID:           int(dbOrder.ID),
		CustomerName: dbOrder.CustomerName,
		OrderedAt:    dbOrder.OrderedAt.Time,
		Items:        items,
	}, nil
}

func (r *orderRepo) GetOrders(ctx context.Context) ([]order.Order, error) {
	dbOrders, err := r.Queries.GetOrders(ctx)
	if err != nil {
		return []order.Order{}, err
	}

	orders := []order.Order{}
	if len(dbOrders) == 0 {
		return orders, nil
	}

	for _, dbOrder := range dbOrders {
		dbItems, err := r.Queries.GetItemByOrderID(ctx, dbOrder.ID)
		if err != nil {
			return []order.Order{}, err
		}

		items := []order.Item{}
		for _, dbItem := range dbItems {
			items = append(items, order.Item{
				ID:          int(dbItem.ID),
				Code:        dbItem.Code,
				Description: dbItem.Description,
				Quantity:    int(dbItem.Quantity),
				OrderID:     int(dbItem.OrderID),
			})
		}

		orders = append(orders, order.Order{
			ID:           int(dbOrder.ID),
			CustomerName: dbOrder.CustomerName,
			OrderedAt:    dbOrder.OrderedAt.Time,
			Items:        items,
		})
	}
	return orders, nil
}

func (r *orderRepo) UpdateOrder(ctx context.Context, o order.Order) (order.Order, error) {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return order.Order{}, err
	}
	defer tx.Rollback(ctx)

	// update order
	_, err = r.Queries.UpdateOrder(ctx, db.UpdateOrderParams{
		ID:           int32(o.ID),
		CustomerName: o.CustomerName,
	})
	if err != nil {
		return order.Order{}, err
	}

	// delete all items
	err = r.Queries.DeleteItemByOrderID(ctx, int32(o.ID))
	if err != nil {
		return order.Order{}, err
	}

	// insert new items
	for _, item := range o.Items {
		_, err := r.Queries.CreateItem(ctx, db.CreateItemParams{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    int32(item.Quantity),
			OrderID:     int32(o.ID),
		})
		if err != nil {
			return order.Order{}, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return order.Order{}, err
	}

	return r.GetOrderByID(ctx, o.ID)
}

func (r *orderRepo) DeleteOrder(ctx context.Context, orderID int) error {
	return r.Queries.DeleteOrder(ctx, int32(orderID))
}
