package main

import (
	"fmt"

	"cafe-system/adapter"
	"cafe-system/decorator"
)

func main() {

	// -------- DECORATOR --------

	var drink decorator.Beverage = &decorator.Espresso{}

	drink = &decorator.Milk{decorator.BeverageDecorator{drink}}
	drink = &decorator.Sugar{decorator.BeverageDecorator{drink}}
	drink = &decorator.WhippedCream{decorator.BeverageDecorator{drink}}

	fmt.Println("Drink:", drink.GetDescription())
	fmt.Println("Cost:", drink.Cost())

	fmt.Println()

	// другой напиток

	var drink2 decorator.Beverage = &decorator.Latte{}
	drink2 = &decorator.Caramel{decorator.BeverageDecorator{drink2}}
	drink2 = &decorator.Milk{decorator.BeverageDecorator{drink2}}

	fmt.Println("Drink:", drink2.GetDescription())
	fmt.Println("Cost:", drink2.Cost())

	fmt.Println()

	// -------- ADAPTER --------

	var processor adapter.IPaymentProcessor

	// PayPal
	processor = &adapter.PayPalPaymentProcessor{}
	processor.ProcessPayment(25.50)

	// Stripe
	stripeService := &adapter.StripePaymentService{}
	processor = adapter.NewStripeAdapter(stripeService)
	processor.ProcessPayment(40.00)

	// Square
	squareService := &adapter.SquarePaymentService{}
	processor = adapter.NewSquareAdapter(squareService)
	processor.ProcessPayment(55.00)
}