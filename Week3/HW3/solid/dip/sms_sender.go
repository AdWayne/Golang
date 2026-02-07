package dip

import "fmt"

type SmsSender struct{}

func (s SmsSender) Send(message string) {
	fmt.Println("SMS sent:", message)
}
