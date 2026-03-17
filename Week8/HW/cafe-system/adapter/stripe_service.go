package adapter

import "fmt"

type StripePaymentService struct{}

func (s *StripePaymentService) MakeTransaction(totalAmount float64) {
	fmt.Printf("Transaction of $%.2f completed via Stripe\n", totalAmount)
}