package payment

// IPaymentStrategy — интерфейс стратегии
type IPaymentStrategy interface {
	Pay(amount float64)
}