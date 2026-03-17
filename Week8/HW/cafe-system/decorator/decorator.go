package decorator

type BeverageDecorator struct {
	Beverage Beverage
}

func (d *BeverageDecorator) GetDescription() string {
	return d.Beverage.GetDescription()
}

func (d *BeverageDecorator) Cost() float64 {
	return d.Beverage.Cost()
}