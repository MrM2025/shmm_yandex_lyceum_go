package main

import (
	"log"
)

func NewOrderLogger() *OrderLogger{
    return &OrderLogger{}
}

type Order struct {
	OrderNumber int
	CustomerName string
	OrderAmount float64
}

func (logger *OrderLogger) AddOrder(order Order) {
	order := Order{logger}
	log.Println("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
}