package decorator

type Milk struct {
	BeverageDecorator
}

func (m *Milk) GetDescription() string {
	return m.Beverage.GetDescription() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + 0.5
}