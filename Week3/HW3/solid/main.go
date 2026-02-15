package main

import (
	"solid/dip"
	"solid/ocp"
	"solid/srp"
)

func main() {
	// SRP
	order := srp.Order{"Book", 2, 500}
	priceCalc := srp.PriceCalculator{}
	price := priceCalc.Calculate(order)

	payment := srp.PaymentProcessor{}
	payment.Process("VISA")

	notify := srp.NotificationService{}
	notify.Send("user@mail.com")

	_ = price

	// OCP
	employee := ocp.Employee{
		Name:       "Alex",
		BaseSalary: 1000,
		Calculator: ocp.PermanentSalary{},
	}
	_ = employee.Calculator.Calculate(employee.BaseSalary)

	// DIP
	service := dip.NewNotificationService(dip.EmailSender{})
	service.Notify("Hello!")
}
