package payment

import "fmt"

type CryptoPayment struct {
	WalletAddress string
}

func (c *CryptoPayment) Pay(amount float64) {
	fmt.Printf("Оплата %.2f тенге криптовалютой (кошелек %s)\n", amount, c.WalletAddress)
}