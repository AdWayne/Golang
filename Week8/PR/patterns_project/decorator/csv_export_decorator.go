package decorator

import "fmt"

type CsvExportDecorator struct {
	*ReportDecorator
}

func NewCsvExportDecorator(report IReport) *CsvExportDecorator {
	return &CsvExportDecorator{
		ReportDecorator: NewReportDecorator(report),
	}
}

func (c *CsvExportDecorator) Generate() string {
	base := c.report.Generate()
	return fmt.Sprintf("%s[Exported to CSV]\n", base)
}