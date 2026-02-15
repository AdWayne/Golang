package srp

import "fmt"

type NotificationService struct{}

func (n NotificationService) Send(email string) {
	fmt.Println("Confirmation email sent to:", email)
}
