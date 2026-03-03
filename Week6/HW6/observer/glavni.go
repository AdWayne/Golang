package main

import (
	"fmt"
	"observer/exchange"
)

func main() {
	exchangeOffice := exchange.CurrencyExchange{}

	mobile := &exchange.MobileApp{Name: "Kaspi App"}
	web := &exchange.WebApp{Site: "Finance.kz"}
	trader := &exchange.Trader{ID: "TR123"}

	exchangeOffice.Register(mobile)
	exchangeOffice.Register(web)
	exchangeOffice.Register(trader)

	fmt.Println("Изменение курса на 480")
	exchangeOffice.SetRate(480)

	fmt.Println("\nУдаляем WebApp\n")
	exchangeOffice.Remove(web)

	fmt.Println("Изменение курса на 520")
	exchangeOffice.SetRate(520)
}