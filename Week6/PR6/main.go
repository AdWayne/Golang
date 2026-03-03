package main

import (
	"fmt"
	"PR6/exchange"
	"PR6/travel"
)

func main() {

	// ---------- STRATEGY ----------
	fmt.Println("=== Travel Booking ===")

	ctx := travel.BookingContext{}
	ctx.SetStrategy(travel.PlaneStrategy{})

	req := travel.TripRequest{
		DistanceKm: 1000,
		Passengers: 2,
		Class:      travel.Business,
		DiscountPct: 10,
	}

	price, _ := ctx.Calculate(req)
	fmt.Println("Price:", price)

	// ---------- OBSERVER ----------
	fmt.Println("\n=== Stock Exchange ===")

	ex := exchange.NewExchange()

	trader := exchange.NewTrader("Alice")
	bot := exchange.NewBot("R2D2", 90, 120)

	ex.Subscribe("AAPL", trader)
	ex.Subscribe("AAPL", bot)

	ex.UpdatePrice("AAPL", 85)
	ex.UpdatePrice("AAPL", 130)
}