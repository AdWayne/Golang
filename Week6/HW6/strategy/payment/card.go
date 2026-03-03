package payment

import "fmt"

type CardPayment struct {
	CardNumber string
}

func (c *CardPayment) Pay(amount float64) {
	fmt.Printf("Оплата %.2f тенге банковской картой %s\n", amount, c.CardNumber)
}