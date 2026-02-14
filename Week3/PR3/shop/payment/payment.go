package payment

import "fmt"
//интерфейс
type Payment interface {
	ProcessPayment(amount float64) error
}

type CreditCardPayment struct{}

func (c CreditCardPayment) ProcessPayment(amount float64) error {
	fmt.Println("Оплата картой:", amount)
	return nil
}
//O open closed жаңа қосу
type PayPalPayment struct{}

func (p PayPalPayment) ProcessPayment(amount float64) error {
	fmt.Println("Оплата PayPal:", amount)
	return nil
}

type BankTransferPayment struct{}

func (b BankTransferPayment) ProcessPayment(amount float64) error {
	fmt.Println("Банковский перевод:", amount)
	return nil
}
