package srp

type PriceCalculator struct{}

func (p PriceCalculator) Calculate(order Order) float64 {
	return float64(order.Quantity) * order.Price * 0.9
}
