package adapter

import "fmt"

type ExternalLogisticsServiceC struct{}

func NewExternalLogisticsServiceC() *ExternalLogisticsServiceC {
	return &ExternalLogisticsServiceC{}
}

func (e *ExternalLogisticsServiceC) StartDelivery(parcelCode string) string {
	return fmt.Sprintf("External C: parcel %s accepted for delivery.", parcelCode)
}

func (e *ExternalLogisticsServiceC) ParcelInfo(parcelCode string) string {
	return fmt.Sprintf("External C: parcel %s status: Processing at warehouse.", parcelCode)
}

func (e *ExternalLogisticsServiceC) DeliveryCost(weight float64) float64 {
	return weight * 5.0
}