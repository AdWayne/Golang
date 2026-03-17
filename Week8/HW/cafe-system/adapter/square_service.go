package adapter

import "fmt"

type SquarePaymentService struct{}

func (s *SquarePaymentService) ExecutePayment(sum float64) {
	fmt.Printf("Payment of $%.2f processed via Square\n", sum)
}