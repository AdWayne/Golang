package payment

type PaymentContext struct {
	strategy IPaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy IPaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) ExecutePayment(amount float64) {
	p.strategy.Pay(amount)
}