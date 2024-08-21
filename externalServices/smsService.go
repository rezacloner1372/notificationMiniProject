package externalServices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
	// TODO
}

func (e *SmsService) SendSms(order *entities.Order) {
	// TODO: Implement the logic to send email
	fmt.Printf("Sms Sent : %v\n", order)

}

func NewSmsService() *SmsService {
	return &SmsService{}
}