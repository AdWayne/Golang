package adapter

import "fmt"

type ExternalLogisticsServiceA struct{}

func NewExternalLogisticsServiceA() *ExternalLogisticsServiceA {
	return &ExternalLogisticsServiceA{}
}

func (e *ExternalLogisticsServiceA) ShipItem(itemId int) string {
	return fmt.Sprintf("External A: item %d shipped.", itemId)
}

func (e *ExternalLogisticsServiceA) TrackShipment(shipmentId int) string {
	return fmt.Sprintf("External A: shipment %d status: On the way.", shipmentId)
}

func (e *ExternalLogisticsServiceA) GetShippingPrice(weight float64) float64 {
	return weight * 3.2
}