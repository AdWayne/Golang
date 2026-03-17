package decorator

type Sugar struct {
	BeverageDecorator
}

func (s *Sugar) GetDescription() string {
	return s.Beverage.GetDescription() + ", Sugar"
}

func (s *Sugar) Cost() float64 {
	return s.Beverage.Cost() + 0.2
}