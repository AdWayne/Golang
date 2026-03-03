package exchange

type Subject interface {
	Subscribe(ticker string, obs Observer)
	Unsubscribe(ticker string, obsID string)
	UpdatePrice(ticker string, price float64)
}