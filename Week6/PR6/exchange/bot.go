package exchange

import "fmt"

type Bot struct {
	id        string
	buyBelow  float64
	sellAbove float64
}

func NewBot(id string, buy, sell float64) *Bot {
	return &Bot{id: id, buyBelow: buy, sellAbove: sell}
}

func (b *Bot) ID() string {
	return b.id
}

func (b *Bot) Notify(ticker string, price float64) {
	if price <= b.buyBelow {
		fmt.Println("Bot", b.id, "BUY", ticker)
	} else if price >= b.sellAbove {
		fmt.Println("Bot", b.id, "SELL", ticker)
	} else {
		fmt.Println("Bot", b.id, "HOLD", ticker)
	}
}