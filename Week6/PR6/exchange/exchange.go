package exchange

import "fmt"

type StockExchange struct {
	prices map[string]float64
	subs   map[string]map[string]Observer
}

func NewExchange() *StockExchange {
	return &StockExchange{
		prices: make(map[string]float64),
		subs:   make(map[string]map[string]Observer),
	}
}

func (e *StockExchange) Subscribe(ticker string, obs Observer) {
	if e.subs[ticker] == nil {
		e.subs[ticker] = make(map[string]Observer)
	}
	e.subs[ticker][obs.ID()] = obs
	fmt.Println("Subscribed:", obs.ID(), "to", ticker)
}

func (e *StockExchange) Unsubscribe(ticker string, obsID string) {
	delete(e.subs[ticker], obsID)
	fmt.Println("Unsubscribed:", obsID, "from", ticker)
}

func (e *StockExchange) UpdatePrice(ticker string, price float64) {
	e.prices[ticker] = price
	fmt.Println("Price updated:", ticker, price)

	for _, obs := range e.subs[ticker] {
		go obs.Notify(ticker, price)
	}
}