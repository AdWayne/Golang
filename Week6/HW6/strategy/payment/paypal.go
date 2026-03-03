package payment

import "fmt"

type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Оплата %.2f тенге через PayPal (%s)\n", amount, p.Email)
}