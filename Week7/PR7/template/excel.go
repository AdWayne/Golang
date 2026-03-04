package main

import "fmt"

type ExcelReport struct{}

func (e *ExcelReport) Header() {
	fmt.Println("Excel Header")
}

func (e *ExcelReport) Body() {
	fmt.Println("Excel Body")
}

func (e *ExcelReport) Footer() {
	fmt.Println("Excel Footer")
}