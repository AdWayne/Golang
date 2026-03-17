package adapter

import "fmt"

type LogisticsAdapterC struct {
	service *ExternalLogisticsServiceC
}

func NewLogisticsAdapterC(service *ExternalLogisticsServiceC) *LogisticsAdapterC {
	return &LogisticsAdapterC{service: service}
}

func (c *LogisticsAdapterC) DeliverOrder(orderId string) string {
	Log("AdapterC: starting DeliverOrder")
	return c.service.StartDelivery("PARCEL-" + orderId)
}

func (c *LogisticsAdapterC) GetDeliveryStatus(orderId string) string {
	Log("AdapterC: getting status")
	return c.service.ParcelInfo("PARCEL-" + orderId)
}

func (c *LogisticsAdapterC) CalculateDeliveryCost(weight float64) float64 {
	Log(fmt.Sprintf("AdapterC: calculating cost for weight %.2f", weight))
	return c.service.DeliveryCost(weight)
}