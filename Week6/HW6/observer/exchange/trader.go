package exchange

import "fmt"

type Trader struct {
	ID string
}

func (t *Trader) Update(rate float64) {
	if rate > 500 {
		fmt.Printf("💰 Трейдер %s продает валюту по курсу %.2f\n", t.ID, rate)
	} else {
		fmt.Printf("📉 Трейдер %s ожидает роста (курс %.2f)\n", t.ID, rate)
	}
}