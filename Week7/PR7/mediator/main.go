package main

func main() {
	chat := &ChatRoom{}

	user1 := NewUser("Alice", chat)
	user2 := NewUser("Bob", chat)

	chat.Register(user1)
	chat.Register(user2)

	user1.Send("Hello Bob!")
	user2.Send("Hi Alice!")
}