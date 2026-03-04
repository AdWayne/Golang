package main

type Mediator interface {
	SendMessage(msg string, user *User)
	Register(user *User)
}