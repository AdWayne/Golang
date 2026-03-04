package mediator

import "fmt"

type User struct {
	Name string
	room Mediator
}

func NewUser(name string, room Mediator) *User {
	return &User{name, room}
}

func (u *User) Send(msg string) {
	u.room.SendMessage(msg, u)
}

func (u *User) Receive(msg string, sender string) {
	fmt.Println(sender, "->", u.Name, ":", msg)
}