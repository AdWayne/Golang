package notification

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("EMAIL:", message)
}

type SmsNotification struct{}

func (s SmsNotification) Send(message string) {
	fmt.Println("SMS:", message)
}
