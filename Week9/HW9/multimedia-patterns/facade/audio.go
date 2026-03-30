package facade

import "fmt"

type AudioSystem struct{}

func (a *AudioSystem) On() {
	fmt.Println("Аудиосистема включена")
}

func (a *AudioSystem) Off() {
	fmt.Println("Аудиосистема выключена")
}

func (a *AudioSystem) SetVolume(v int) {
	fmt.Println("Громкость:", v)
}