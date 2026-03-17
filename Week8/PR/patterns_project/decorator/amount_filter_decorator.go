package decorator

import "fmt"

type AmountFilterDecorator struct {
	*ReportDecorator
	minAmount float64
}

func NewAmountFilterDecorator(report IReport, minAmount float64) *AmountFilterDecorator {
	return &AmountFilterDecorator{
		ReportDecorator: NewReportDecorator(report),
		minAmount:       minAmount,
	}
}

func (a *AmountFilterDecorator) Generate() string {
	base := a.report.Generate()
	return fmt.Sprintf("%s[Sales Amount Filter Applied: amount >= %.2f]\n", base, a.minAmount)
}