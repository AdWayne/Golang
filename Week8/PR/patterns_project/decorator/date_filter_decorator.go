package decorator

import "fmt"

type DateFilterDecorator struct {
	*ReportDecorator
	from string
	to   string
}

func NewDateFilterDecorator(report IReport, from, to string) *DateFilterDecorator {
	return &DateFilterDecorator{
		ReportDecorator: NewReportDecorator(report),
		from:            from,
		to:              to,
	}
}

func (d *DateFilterDecorator) Generate() string {
	base := d.report.Generate()
	return fmt.Sprintf("%s[Date Filter Applied: from %s to %s]\n", base, d.from, d.to)
}