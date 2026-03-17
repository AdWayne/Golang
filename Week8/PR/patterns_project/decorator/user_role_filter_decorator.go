package decorator

import "fmt"

type UserRoleFilterDecorator struct {
	*ReportDecorator
	role string
}

func NewUserRoleFilterDecorator(report IReport, role string) *UserRoleFilterDecorator {
	return &UserRoleFilterDecorator{
		ReportDecorator: NewReportDecorator(report),
		role:            role,
	}
}

func (u *UserRoleFilterDecorator) Generate() string {
	base := u.report.Generate()
	return fmt.Sprintf("%s[User Role Filter Applied: role = %s]\n", base, u.role)
}