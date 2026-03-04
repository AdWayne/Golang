package command

import "fmt"

type Invoker struct {
	history []Command
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) Execute(cmd Command) {
	cmd.Execute()
	i.history = append(i.history, cmd)
}

func (i *Invoker) Undo() {
	if len(i.history) == 0 {
		fmt.Println("Nothing to undo")
		return
	}

	last := i.history[len(i.history)-1]
	last.Undo()

	i.history = i.history[:len(i.history)-1]
}