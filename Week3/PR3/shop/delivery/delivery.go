package delivery

import (
	"fmt"
	"shop/order"
)

type Delivery interface {
	Deliver(order *order.Order) error
}

type CourierDelivery struct{}

func (c CourierDelivery) Deliver(order *order.Order) error {
	fmt.Println("Курьер доставляет заказ №", order.ID)
	return nil
}

type PostDelivery struct{}

func (p PostDelivery) Deliver(order *order.Order) error {
	fmt.Println("Почта доставляет заказ №", order.ID)
	return nil
}

type PickUpPointDelivery struct{}

func (p PickUpPointDelivery) Deliver(order *order.Order) error {
	fmt.Println("Самовывоз заказа №", order.ID)
	return nil
}