package externalServices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
	// TODO
}

func (e *SmsService) SendSms(order *entities.Order) {
	fmt.Printf("Sms Sent : %v\n", order)

}

func (e *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("Sms Sent : receiver %s\n, message %s\n", receiver, message)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
