package decorator

func BuildReport(reportType string, decorators []string) IReport {
	var report IReport

	switch reportType {
	case "sales":
		report = NewSalesReport()
	case "users":
		report = NewUserReport()
	default:
		report = NewSalesReport()
	}

	for _, dec := range decorators {
		switch dec {
		case "datefilter":
			report = NewDateFilterDecorator(report, "2025-01-01", "2025-12-31")
		case "sort":
			report = NewSortingDecorator(report, "date")
		case "csv":
			report = NewCsvExportDecorator(report)
		case "pdf":
			report = NewPdfExportDecorator(report)
		case "amountfilter":
			report = NewAmountFilterDecorator(report, 1000)
		case "rolefilter":
			report = NewUserRoleFilterDecorator(report, "premium")
		}
	}

	return report
}