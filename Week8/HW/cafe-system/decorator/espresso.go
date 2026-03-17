package decorator

type Espresso struct{}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) Cost() float64 {
	return 2.0
}