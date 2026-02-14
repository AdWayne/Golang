package main
//SRP a lot of struct
//ISP a lot of interface
//DIP in interface
//OCP no change code
//LSP no change interface

import (
	"shop/delivery"
	"shop/discount"
	"shop/notification"
	"shop/order"
	"shop/payment"
	"shop/service"
)

func main() {
	o := order.Order{ID: 1}

	o.AddItem(order.Item{Name: "Ноутбук", Price: 300000, Qty: 1})
	o.AddItem(order.Item{Name: "Мышь", Price: 5000, Qty: 2})
	o.AddItem(order.Item{Name: "Клавиатура", Price: 5000, Qty: 2})

	o.Discounts = append(o.Discounts,
		discount.PercentageDiscount{Percent: 10},
		discount.FixedDiscount{Value: 5000},
	)

	// p := payment.CreditCardPayment{}
	// d := delivery.CourierDelivery{}
	// n := notification.EmailNotification{}

	p := payment.PayPalPayment{}
	d := delivery.PostDelivery{}
	n := notification.SmsNotification{}


	orderService := service.NewOrderService(p, d, n)
	orderService.Checkout(&o)
}