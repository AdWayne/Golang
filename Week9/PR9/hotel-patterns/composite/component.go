package composite

type OrganizationComponent interface {
	Show(indent string)
	GetSalary() float64
	GetEmployeeCount() int
}