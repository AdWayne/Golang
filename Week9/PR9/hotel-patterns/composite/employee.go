package composite

import "fmt"

type Employee struct {
	Name     string
	Position string
	Salary   float64
}

func (e *Employee) Show(indent string) {
	fmt.Println(indent, e.Name, "-", e.Position)
}

func (e *Employee) GetSalary() float64 {
	return e.Salary
}

func (e *Employee) GetEmployeeCount() int {
	return 1
}