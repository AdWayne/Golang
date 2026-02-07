package service

import (
	"shop/delivery"
	"shop/notification"
	"shop/order"
	"shop/payment"
)

type OrderService struct {
	payment      payment.Payment
	delivery     delivery.Delivery
	notification notification.Notification
}

func NewOrderService(
	p payment.Payment,
	d delivery.Delivery,
	n notification.Notification,
) *OrderService {
	return &OrderService{
		payment:      p,
		delivery:     d,
		notification: n,
	}
}

func (s *OrderService) Checkout(order *order.Order) {
	total := order.CalculateTotal()

	s.payment.ProcessPayment(total)
	s.delivery.Deliver(order)
	s.notification.Send("Заказ успешно оформлен!")
}
