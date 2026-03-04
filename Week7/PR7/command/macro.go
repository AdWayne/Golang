package main

import "fmt"

type MacroCommand struct {
	commands []Command
}

func (m *MacroCommand) Execute() error {
	fmt.Println("Macro START")
	for _, c := range m.commands {
		c.Execute()
	}
	fmt.Println("Macro END")
	return nil
}

func (m *MacroCommand) Undo() error {
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
	return nil
}

func (m *MacroCommand) Name() string {
	return "MacroCommand"
}