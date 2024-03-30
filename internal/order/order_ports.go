package order

type Repository interface {
	Create(order *Order) (*Order, error)
	Get() (*[]Order, error)
	Update(order *Order) (*Order, error)
	Delete(orderId uint) error
}

type Service interface {
	CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrder() (*[]GetOrderResponse, error)
	UpdateOrder(orderId uint, req *UpdateOrderRequest) (*UpdateOrderResponse, error)
	DeleteOrder(req *DeleteOrderRequest) error
}
