package decorator

import "fmt"

type SalesReport struct{}

func NewSalesReport() *SalesReport {
	return &SalesReport{}
}

func (s *SalesReport) Generate() string {
	return fmt.Sprintf(
		"Sales Report\n"+
			"---------------------------------\n"+
			"OrderID: 1001 | Date: 2025-01-10 | Amount: 1500\n"+
			"OrderID: 1002 | Date: 2025-03-15 | Amount: 700\n"+
			"OrderID: 1003 | Date: 2025-05-20 | Amount: 2300\n",
	)
}