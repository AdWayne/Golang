package composite

import "fmt"

type Department struct {
	Name       string
	components []OrganizationComponent
}

func NewDepartment(name string) *Department {
	return &Department{Name: name}
}

func (d *Department) Add(c OrganizationComponent) {
	d.components = append(d.components, c)
}

func (d *Department) Show(indent string) {
	fmt.Println(indent, "Department:", d.Name)
	for _, c := range d.components {
		c.Show(indent + "  ")
	}
}

func (d *Department) GetSalary() float64 {
	total := 0.0
	for _, c := range d.components {
		total += c.GetSalary()
	}
	return total
}

func (d *Department) GetEmployeeCount() int {
	count := 0
	for _, c := range d.components {
		count += c.GetEmployeeCount()
	}
	return count
}