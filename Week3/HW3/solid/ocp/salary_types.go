package ocp

type PermanentSalary struct{}

func (p PermanentSalary) Calculate(baseSalary float64) float64 {
	return baseSalary * 1.2
}

type ContractSalary struct{}

func (c ContractSalary) Calculate(baseSalary float64) float64 {
	return baseSalary * 1.1
}

type InternSalary struct{}

func (i InternSalary) Calculate(baseSalary float64) float64 {
	return baseSalary * 0.8
}
