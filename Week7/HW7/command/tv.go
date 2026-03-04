package command

import "fmt"

type TV struct{}

func (tv *TV) On() {
	fmt.Println("TV ON")
}

func (tv *TV) Off() {
	fmt.Println("TV OFF")
}