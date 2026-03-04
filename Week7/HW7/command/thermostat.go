package command

import "fmt"

type Thermostat struct {
	Temperature int
}

func (t *Thermostat) Increase() {
	t.Temperature++
	fmt.Println("Temperature:", t.Temperature)
}

func (t *Thermostat) Decrease() {
	t.Temperature--
	fmt.Println("Temperature:", t.Temperature)
}