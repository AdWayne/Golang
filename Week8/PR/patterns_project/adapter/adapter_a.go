package adapter

import (
	"fmt"
	"strconv"
)

type LogisticsAdapterA struct {
	service *ExternalLogisticsServiceA
}

func NewLogisticsAdapterA(service *ExternalLogisticsServiceA) *LogisticsAdapterA {
	return &LogisticsAdapterA{service: service}
}

func (a *LogisticsAdapterA) DeliverOrder(orderId string) string {
	Log("AdapterA: starting DeliverOrder")

	id, err := strconv.Atoi(orderId)
	if err != nil {
		Log("AdapterA error: invalid order ID")
		return "AdapterA error: invalid order ID"
	}

	result := a.service.ShipItem(id)
	Log("AdapterA: order delivered successfully")
	return result
}

func (a *LogisticsAdapterA) GetDeliveryStatus(orderId string) string {
	Log("AdapterA: getting status")

	id, err := strconv.Atoi(orderId)
	if err != nil {
		Log("AdapterA error: invalid shipment ID")
		return "AdapterA error: invalid shipment ID"
	}

	return a.service.TrackShipment(id)
}

func (a *LogisticsAdapterA) CalculateDeliveryCost(weight float64) float64 {
	Log(fmt.Sprintf("AdapterA: calculating cost for weight %.2f", weight))
	return a.service.GetShippingPrice(weight)
}