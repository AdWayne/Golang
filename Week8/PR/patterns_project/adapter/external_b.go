package adapter

import "fmt"

type ExternalLogisticsServiceB struct{}

func NewExternalLogisticsServiceB() *ExternalLogisticsServiceB {
	return &ExternalLogisticsServiceB{}
}

func (e *ExternalLogisticsServiceB) SendPackage(packageInfo string) string {
	return fmt.Sprintf("External B: package [%s] sent.", packageInfo)
}

func (e *ExternalLogisticsServiceB) CheckPackageStatus(trackingCode string) string {
	return fmt.Sprintf("External B: tracking code %s status: Delivered.", trackingCode)
}

func (e *ExternalLogisticsServiceB) EstimateCost(weight float64) float64 {
	return weight * 4.1
}