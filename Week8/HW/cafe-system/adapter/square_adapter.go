package adapter

type SquarePaymentAdapter struct {
	square *SquarePaymentService
}

func NewSquareAdapter(service *SquarePaymentService) *SquarePaymentAdapter {
	return &SquarePaymentAdapter{square: service}
}

func (s *SquarePaymentAdapter) ProcessPayment(amount float64) {
	s.square.ExecutePayment(amount)
}