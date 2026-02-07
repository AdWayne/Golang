package discount

type PercentageDiscount struct {
	Percent float64
}

func (p PercentageDiscount) Apply(amount float64) float64 {
	return amount * (1 - p.Percent/100)
}

type FixedDiscount struct {
	Value float64
}

func (f FixedDiscount) Apply(amount float64) float64 {
	return amount - f.Value
}
