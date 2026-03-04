package main

import "fmt"

type LightOn struct{ l *Light }

func (c LightOn) Execute() error { c.l.On(); return nil }
func (c LightOn) Undo() error    { c.l.Off(); return nil }
func (c LightOn) Name() string   { return "LightOn" }

type TVOn struct{ t *TV }

func (c TVOn) Execute() error { c.t.On(); return nil }
func (c TVOn) Undo() error    { c.t.Off(); return nil }
func (c TVOn) Name() string   { return "TVOn" }

func main() {
	light := &Light{"Living Room", false}
	tv := &TV{"Bedroom", false}

	remote := NewRemote(2)

	remote.Set(0, LightOn{light})
	remote.Set(1, TVOn{tv}) // <-- используем tv

	remote.Press(0)
	remote.Press(1)

	remote.Undo()
	remote.Redo()

	macro := &MacroCommand{
		commands: []Command{
			LightOn{light},
			TVOn{tv},
		},
	}

	macro.Execute()

	fmt.Println("Command pattern finished")
}