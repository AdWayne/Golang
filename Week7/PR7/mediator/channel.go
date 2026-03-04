package main

type ChatRoom struct {
	users []*User
}

func (c *ChatRoom) Register(user *User) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) SendMessage(msg string, sender *User) {
	for _, u := range c.users {
		if u != sender {
			u.Receive(sender.name + ": " + msg)
		}
	}
}