package command

import "fmt"

type Door struct {
}

func (d *Door) Open() {
	fmt.Println("Door OPEN")
}

func (d *Door) Close() {
	fmt.Println("Door CLOSE")
}