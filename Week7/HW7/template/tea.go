package template

import "fmt"

type Tea struct {
	BaseBeverage
}

func (t Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}

func (t Tea) Prepare() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}