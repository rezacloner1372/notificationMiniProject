package externalServices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct {
	// TODO
}

// a method that get order and send email
func (e *EmailService) SendEmail(order *entities.Order) {
	// TODO: Implement the logic to send email
	fmt.Printf("Email Sent : %v\n", order)

}

func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("email Sent : receiver %s\n, message %s\n", receiver, message)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
