package main

import (
	"notification/entities"
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

	orderService := services.NewOrderService()

	orderService.CreateOrder(&order1)
}
