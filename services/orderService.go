package services

import (
	"fmt"
	"notification/entities"
	"notification/externalServices"
)

// register the order
type OrderService struct {
	EmaileService *externalServices.EmailService
	SmsService    * externalServices.SmsService
}

// Method to create a new order
func (o *OrderService) CreateOrder(order *entities.Order) (*entities.Order) {
	// TODO: Implement the logic to register the order
	fmt.Printf("Order Created : %v\n", order)
	// send sms or email notifiaction
	o.EmaileService.SendEmail(order)
	o.SmsService.SendSms(order)

	return order
}

// create method New for OrderService
func NewOrderService() *OrderService {
	return &OrderService{
		EmaileService: externalServices.NewEmailService(),
		SmsService: externalServices.NewSmsService(),
	}
	
}