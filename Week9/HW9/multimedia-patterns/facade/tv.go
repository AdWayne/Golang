package facade

import "fmt"

type TV struct{}

func (t *TV) On() {
	fmt.Println("TV включен")
}

func (t *TV) Off() {
	fmt.Println("TV выключен")
}

func (t *TV) SetChannel(channel int) {
	fmt.Println("TV канал:", channel)
}