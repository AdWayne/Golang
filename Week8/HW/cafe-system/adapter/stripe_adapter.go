package adapter

type StripePaymentAdapter struct {
	stripe *StripePaymentService
}

func NewStripeAdapter(service *StripePaymentService) *StripePaymentAdapter {
	return &StripePaymentAdapter{stripe: service}
}

func (s *StripePaymentAdapter) ProcessPayment(amount float64) {
	s.stripe.MakeTransaction(amount)
}