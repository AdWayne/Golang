package srp

import "fmt"

type PaymentProcessor struct{}

func (p PaymentProcessor) Process(details string) {
	fmt.Println("Payment processed using:", details)
}
