package main

import (
	"fmt"
	"strategy/payment"
)

func main() {
	var choice int
	var amount float64

	fmt.Println("Введите сумму оплаты:")
	fmt.Scanln(&amount)

	fmt.Println("Выберите способ оплаты:")
	fmt.Println("1 - Банковская карта")
	fmt.Println("2 - PayPal")
	fmt.Println("3 - Криптовалюта")
	fmt.Scanln(&choice)

	context := payment.PaymentContext{}

	switch choice {
	case 1:
		context.SetStrategy(&payment.CardPayment{CardNumber: "1234-5678-0000-9999"})
	case 2:
		context.SetStrategy(&payment.PayPalPayment{Email: "user@gmail.com"})
	case 3:
		context.SetStrategy(&payment.CryptoPayment{WalletAddress: "0xABCDEF123456"})
	default:
		fmt.Println("Неверный выбор")
		return
	}

	context.ExecutePayment(amount)
}