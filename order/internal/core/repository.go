package core

//secondary port

type OrderRepository interface {
	Save(order *Order) error
	FindByID(id string) (*Order, error)
	UpdateOrderStatusByID(orderID string, status string) error
}
