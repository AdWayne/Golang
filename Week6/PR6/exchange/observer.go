package exchange

type Observer interface {
	ID() string
	Notify(ticker string, price float64)
}