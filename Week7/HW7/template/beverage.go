package template

import "fmt"

type Beverage interface {
	Brew()
	AddCondiments()
}

type BaseBeverage struct{}

func (b *BaseBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (b *BaseBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}