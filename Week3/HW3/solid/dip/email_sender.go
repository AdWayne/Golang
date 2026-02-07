package dip

import "fmt"

type EmailSender struct{}

func (e EmailSender) Send(message string) {
	fmt.Println("Email sent:", message)
}
