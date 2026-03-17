package decorator

type Caramel struct {
	BeverageDecorator
}

func (c *Caramel) GetDescription() string {
	return c.Beverage.GetDescription() + ", Caramel"
}

func (c *Caramel) Cost() float64 {
	return c.Beverage.Cost() + 0.6
}