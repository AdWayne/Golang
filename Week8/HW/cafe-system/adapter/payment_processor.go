package adapter

type IPaymentProcessor interface {
	ProcessPayment(amount float64)
}