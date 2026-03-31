package composite

import "fmt"

type Contractor struct {
	Name string
	Pay  float64
}

func (c *Contractor) Show(indent string) {
	fmt.Println(indent, c.Name, "- contractor")
}

func (c *Contractor) GetSalary() float64 {
	return 0
}

func (c *Contractor) GetEmployeeCount() int {
	return 1
}