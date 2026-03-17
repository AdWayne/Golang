package adapter

type DeliveryServiceFactory struct{}

func NewDeliveryServiceFactory() *DeliveryServiceFactory {
	return &DeliveryServiceFactory{}
}

func (f *DeliveryServiceFactory) GetService(serviceType string) IInternalDeliveryService {
	switch serviceType {
	case "internal":
		return NewInternalDeliveryService()
	case "externalA":
		return NewLogisticsAdapterA(NewExternalLogisticsServiceA())
	case "externalB":
		return NewLogisticsAdapterB(NewExternalLogisticsServiceB())
	case "externalC":
		return NewLogisticsAdapterC(NewExternalLogisticsServiceC())
	default:
		return NewInternalDeliveryService()
	}
}