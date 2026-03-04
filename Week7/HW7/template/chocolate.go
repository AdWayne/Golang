package template

import "fmt"

type Chocolate struct {
	BaseBeverage
}

func (h Chocolate) Brew() {
	fmt.Println("Mixing chocolate powder")
}

func (h Chocolate) AddCondiments() {
	fmt.Println("Adding marshmallow")
}

func (h Chocolate) Prepare() {
	h.BoilWater()
	h.Brew()
	h.PourInCup()
	h.AddCondiments()
}