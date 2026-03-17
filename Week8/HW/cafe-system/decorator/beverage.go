package decorator

type Beverage interface {
	GetDescription() string
	Cost() float64
}