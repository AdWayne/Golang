package dip

type NotificationService struct {
	sender MessageSender
}

func NewNotificationService(sender MessageSender) NotificationService {
	return NotificationService{sender: sender}
}

func (n NotificationService) Notify(message string) {
	n.sender.Send(message)
}
