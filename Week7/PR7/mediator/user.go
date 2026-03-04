package main

import "fmt"

type User struct {
	name     string
	mediator Mediator
}

func NewUser(name string, m Mediator) *User {
	return &User{name, m}
}

func (u *User) Send(msg string) {
	u.mediator.SendMessage(msg, u)
}

func (u *User) Receive(msg string) {
	fmt.Println(u.name, "received:", msg)
}