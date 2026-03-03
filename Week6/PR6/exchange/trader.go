package exchange

import "fmt"

type Trader struct {
	id string
}

func NewTrader(id string) *Trader {
	return &Trader{id: id}
}

func (t *Trader) ID() string {
	return t.id
}

func (t *Trader) Notify(ticker string, price float64) {
	fmt.Println("Trader", t.id, "got update:", ticker, price)
}