package decorator

type ReportDecorator struct {
	report IReport
}

func NewReportDecorator(report IReport) *ReportDecorator {
	return &ReportDecorator{report: report}
}

func (d *ReportDecorator) Generate() string {
	return d.report.Generate()
}