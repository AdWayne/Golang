package main

import "fmt"

// Шаги отчета
type ReportSteps interface {
	Header()
	Body()
	Footer()
}

// Шаблонный метод
func GenerateReport(r ReportSteps) {
	fmt.Println("Start report generation")

	r.Header()
	r.Body()
	r.Footer()

	fmt.Println("End report generation")
}