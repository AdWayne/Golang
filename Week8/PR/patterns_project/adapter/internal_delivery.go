package adapter

import "fmt"

type InternalDeliveryService struct{}

func NewInternalDeliveryService() *InternalDeliveryService {
	return &InternalDeliveryService{}
}

func (i *InternalDeliveryService) DeliverOrder(orderId string) string {
	return fmt.Sprintf("Internal delivery: order %s has been delivered.", orderId)
}

func (i *InternalDeliveryService) GetDeliveryStatus(orderId string) string {
	return fmt.Sprintf("Internal delivery status for order %s: In transit.", orderId)
}

func (i *InternalDeliveryService) CalculateDeliveryCost(weight float64) float64 {
	return weight * 2.5
}