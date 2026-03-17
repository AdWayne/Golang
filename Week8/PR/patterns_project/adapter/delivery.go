package adapter

type IInternalDeliveryService interface {
	DeliverOrder(orderId string) string
	GetDeliveryStatus(orderId string) string
	CalculateDeliveryCost(weight float64) float64
}