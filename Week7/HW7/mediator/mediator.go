package mediator

type Mediator interface {
	SendMessage(msg string, user *User)
}