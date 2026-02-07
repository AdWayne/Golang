package isp

import "fmt"

type AllInOnePrinter struct{}

func (a AllInOnePrinter) Print(content string) {
	fmt.Println("Printing:", content)
}

func (a AllInOnePrinter) Scan(content string) {
	fmt.Println("Scanning:", content)
}

func (a AllInOnePrinter) Fax(content string) {
	fmt.Println("Faxing:", content)
}

type BasicPrinter struct{}

func (b BasicPrinter) Print(content string) {
	fmt.Println("Printing:", content)
}
