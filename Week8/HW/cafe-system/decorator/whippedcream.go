package decorator

type WhippedCream struct {
	BeverageDecorator
}

func (w *WhippedCream) GetDescription() string {
	return w.Beverage.GetDescription() + ", Whipped Cream"
}

func (w *WhippedCream) Cost() float64 {
	return w.Beverage.Cost() + 0.7
}