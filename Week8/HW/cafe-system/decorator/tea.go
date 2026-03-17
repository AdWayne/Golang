package decorator

type Tea struct{}

func (t *Tea) GetDescription() string {
	return "Tea"
}

func (t *Tea) Cost() float64 {
	return 1.5
}