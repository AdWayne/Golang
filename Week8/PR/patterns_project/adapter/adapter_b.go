package adapter

import "fmt"

type LogisticsAdapterB struct {
	service *ExternalLogisticsServiceB
}

func NewLogisticsAdapterB(service *ExternalLogisticsServiceB) *LogisticsAdapterB {
	return &LogisticsAdapterB{service: service}
}

func (b *LogisticsAdapterB) DeliverOrder(orderId string) string {
	Log("AdapterB: starting DeliverOrder")
	return b.service.SendPackage("Order#" + orderId)
}

func (b *LogisticsAdapterB) GetDeliveryStatus(orderId string) string {
	Log("AdapterB: getting status")
	return b.service.CheckPackageStatus("TRACK-" + orderId)
}

func (b *LogisticsAdapterB) CalculateDeliveryCost(weight float64) float64 {
	Log(fmt.Sprintf("AdapterB: calculating cost for weight %.2f", weight))
	return b.service.EstimateCost(weight)
}