package main

import (
	"fmt"
	"patterns_project/adapter"
	"patterns_project/decorator"
)

func main() {
	fmt.Println("===================================")
	fmt.Println("ДЕКОРАТОР: СИСТЕМА ОТЧЕТНОСТИ")
	fmt.Println("===================================")

	// Базовый отчет по продажам
	var salesReport decorator.IReport = decorator.NewSalesReport()

	// Применяем несколько декораторов
	salesReport = decorator.NewDateFilterDecorator(salesReport, "2025-01-01", "2025-12-31")
	salesReport = decorator.NewSortingDecorator(salesReport, "date")
	salesReport = decorator.NewAmountFilterDecorator(salesReport, 1000.0)
	salesReport = decorator.NewCsvExportDecorator(salesReport)

	fmt.Println("\nОтчет по продажам с декораторами:")
	fmt.Println(salesReport.Generate())

	// Базовый отчет по пользователям
	var userReport decorator.IReport = decorator.NewUserReport()

	userReport = decorator.NewUserRoleFilterDecorator(userReport, "premium")
	userReport = decorator.NewSortingDecorator(userReport, "name")
	userReport = decorator.NewPdfExportDecorator(userReport)

	fmt.Println("\nОтчет по пользователям с декораторами:")
	fmt.Println(userReport.Generate())

	// Динамический выбор декораторов
	fmt.Println("\nДинамический выбор декораторов:")
	dynamicReport := decorator.BuildReport("sales", []string{"datefilter", "sort", "pdf"})
	fmt.Println(dynamicReport.Generate())

	fmt.Println("\n===================================")
	fmt.Println("АДАПТЕР: СИСТЕМА ЛОГИСТИКИ")
	fmt.Println("===================================")

	internalService := adapter.NewDeliveryServiceFactory().GetService("internal")
	fmt.Println(internalService.DeliverOrder("101"))
	fmt.Println(internalService.GetDeliveryStatus("101"))
	fmt.Printf("Стоимость доставки: %.2f\n", internalService.CalculateDeliveryCost(15.5))

	fmt.Println()

	externalA := adapter.NewDeliveryServiceFactory().GetService("externalA")
	fmt.Println(externalA.DeliverOrder("202"))
	fmt.Println(externalA.GetDeliveryStatus("202"))
	fmt.Printf("Стоимость доставки: %.2f\n", externalA.CalculateDeliveryCost(8.0))

	fmt.Println()

	externalB := adapter.NewDeliveryServiceFactory().GetService("externalB")
	fmt.Println(externalB.DeliverOrder("303"))
	fmt.Println(externalB.GetDeliveryStatus("303"))
	fmt.Printf("Стоимость доставки: %.2f\n", externalB.CalculateDeliveryCost(12.0))

	fmt.Println()

	externalC := adapter.NewDeliveryServiceFactory().GetService("externalC")
	fmt.Println(externalC.DeliverOrder("404"))
	fmt.Println(externalC.GetDeliveryStatus("404"))
	fmt.Printf("Стоимость доставки: %.2f\n", externalC.CalculateDeliveryCost(20.0))
}