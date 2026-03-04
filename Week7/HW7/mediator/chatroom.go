package mediator

import "fmt"

type ChatRoom struct {
	users []*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{}
}

func (c *ChatRoom) Join(user *User) {
	c.users = append(c.users, user)
	fmt.Println(user.Name, "joined chat")
}

func (c *ChatRoom) SendMessage(msg string, sender *User) {

	for _, u := range c.users {

		if u != sender {
			u.Receive(msg, sender.Name)
		}

	}
}