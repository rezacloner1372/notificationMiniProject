package services

import (
	"fmt"
	"notification/entities"
	"notification/externalServices"
)

// register the order
type OrderService struct {
	Notifier externalServices.Notifier
}

// Method to create a new order
func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	// TODO: Implement the logic to register the order
	fmt.Printf("Order Created : %v\n", order)
	// send sms or email notifiaction
	orderService.Notifier.SendNotify(order.UserName, "Order Created")

	return order
}

// create method New for OrderService
func NewOrderService(notifier externalServices.Notifier) *OrderService {
	return &OrderService{
		Notifier: notifier,
	}

}
