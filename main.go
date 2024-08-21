package main

import (
	"notification/entities"
	"notification/externalServices"
	"notification/services"
)

func main() {

	order1 := entities.Order{
		ID:        1,
		UserName:  "MohammadReza Saberi",
		UserEmail: "m.rezasaberi01@gmail.com",
		UserPhone: "09137373384",
		Price:     100,
		Status:    true,
	}

	order2 := entities.Order{
		ID:        2,
		UserName:  "MohammadReza Karimi",
		UserEmail: "karimi@gmail.com",
		UserPhone: "098726378848",
		Price:     120,
		Status:    true,
	}

	orderService := services.NewOrderService(externalServices.NewSmsService())
	orderService.CreateOrder(&order1)

	orderService = services.NewOrderService(externalServices.NewEmailService())
	orderService.CreateOrder(&order2)
}
