package adapter

import "fmt"

type PayPalPaymentProcessor struct{}

func (p *PayPalPaymentProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of $%.2f via PayPal\n", amount)
}