package decorator

type Latte struct{}

func (l *Latte) GetDescription() string {
	return "Latte"
}

func (l *Latte) Cost() float64 {
	return 3.0
}