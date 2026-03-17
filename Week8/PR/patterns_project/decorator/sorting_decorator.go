package decorator

import "fmt"

type SortingDecorator struct {
	*ReportDecorator
	criteria string
}

func NewSortingDecorator(report IReport, criteria string) *SortingDecorator {
	return &SortingDecorator{
		ReportDecorator: NewReportDecorator(report),
		criteria:        criteria,
	}
}

func (s *SortingDecorator) Generate() string {
	base := s.report.Generate()
	return fmt.Sprintf("%s[Sorted By: %s]\n", base, s.criteria)
}