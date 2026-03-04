package main

import "fmt"

type CSVReport struct{}

func (c *CSVReport) Header() {
	fmt.Println("CSV Header")
}

func (c *CSVReport) Body() {
	fmt.Println("CSV Data")
}

func (c *CSVReport) Footer() {
	fmt.Println("CSV Footer")
}