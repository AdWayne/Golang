package facade

import "fmt"

type DVDPlayer struct{}

func (d *DVDPlayer) Play(movie string) {
	fmt.Println("DVD воспроизведение:", movie)
}

func (d *DVDPlayer) Pause() {
	fmt.Println("DVD пауза")
}

func (d *DVDPlayer) Stop() {
	fmt.Println("DVD стоп")
}