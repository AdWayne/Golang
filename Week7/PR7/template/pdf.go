package main

import "fmt"

type PDFReport struct{}

func (p *PDFReport) Header() {
	fmt.Println("PDF Header")
}

func (p *PDFReport) Body() {
	fmt.Println("PDF Body")
}

func (p *PDFReport) Footer() {
	fmt.Println("PDF Footer")
}