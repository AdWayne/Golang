package main

import "fmt"

type HTMLReport struct{}

func (h *HTMLReport) Header() {
	fmt.Println("<h1>HTML Header</h1>")
}

func (h *HTMLReport) Body() {
	fmt.Println("<p>HTML Body</p>")
}

func (h *HTMLReport) Footer() {
	fmt.Println("<footer>HTML Footer</footer>")
}