package main

import (
	"HW7/command"
	"HW7/mediator"
	"HW7/template"
	"fmt"
)

func main() {

	fmt.Println("===== COMMAND PATTERN =====")

	light := &command.Light{}
	invoker := command.NewInvoker()

	lightOn := command.NewLightOnCommand(light)
	lightOff := command.NewLightOffCommand(light)

	invoker.Execute(lightOn)
	invoker.Execute(lightOff)
	invoker.Undo()

	fmt.Println("\n===== TEMPLATE METHOD =====")

	tea := template.Tea{}
	tea.Prepare()

	coffee := template.Coffee{}
	coffee.Prepare()

	fmt.Println("\n===== MEDIATOR =====")

	room := mediator.NewChatRoom()

	alice := mediator.NewUser("Alice", room)
	bob := mediator.NewUser("Bob", room)
	charlie := mediator.NewUser("Charlie", room)

	room.Join(alice)
	room.Join(bob)
	room.Join(charlie)

	alice.Send("Hello everyone!")
}
