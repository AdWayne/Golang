package decorator

import "fmt"

type PdfExportDecorator struct {
	*ReportDecorator
}

func NewPdfExportDecorator(report IReport) *PdfExportDecorator {
	return &PdfExportDecorator{
		ReportDecorator: NewReportDecorator(report),
	}
}

func (p *PdfExportDecorator) Generate() string {
	base := p.report.Generate()
	return fmt.Sprintf("%s[Exported to PDF]\n", base)
}