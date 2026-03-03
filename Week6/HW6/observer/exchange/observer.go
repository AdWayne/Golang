package exchange

type IObserver interface {
	Update(rate float64)
}