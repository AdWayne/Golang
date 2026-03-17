package decorator

type Mocha struct{}

func (m *Mocha) GetDescription() string {
	return "Mocha"
}

func (m *Mocha) Cost() float64 {
	return 3.5
}