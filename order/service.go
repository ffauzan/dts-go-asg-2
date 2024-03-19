package order

import "context"

type service struct {
	repo OrderRepo
}

func NewService(repo OrderRepo) OrderService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateOrder(ctx context.Context, o Order) (Order, error) {
	return s.repo.CreateOrder(ctx, o)
}

func (s *service) GetOrderByID(ctx context.Context, id int) (Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *service) GetOrders(ctx context.Context) ([]Order, error) {
	return s.repo.GetOrders(ctx)
}

func (s *service) UpdateOrder(ctx context.Context, o Order) (Order, error) {
	return s.repo.UpdateOrder(ctx, o)
}

func (s *service) DeleteOrder(ctx context.Context, id int) error {
	return s.repo.DeleteOrder(ctx, id)
}
