package command

import "fmt"

type Light struct {
}

func (l *Light) On() {
	fmt.Println("Light ON")
}

func (l *Light) Off() {
	fmt.Println("Light OFF")
}