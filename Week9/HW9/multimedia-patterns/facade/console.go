package facade

import "fmt"

type GameConsole struct{}

func (g *GameConsole) On() {
	fmt.Println("Консоль включена")
}

func (g *GameConsole) PlayGame(game string) {
	fmt.Println("Запуск игры:", game)
}