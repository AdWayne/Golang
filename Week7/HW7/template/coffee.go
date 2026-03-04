package template

import (
	"bufio"
	"fmt"
	"os"
)

type Coffee struct {
	BaseBeverage
}

func (c Coffee) Brew() {
	fmt.Println("Dripping coffee")
}

func (c Coffee) AddCondiments() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Add milk and sugar? (y/n): ")

	input, _ := reader.ReadString('\n')

	if input == "y\n" {
		fmt.Println("Adding milk and sugar")
	}
}

func (c Coffee) Prepare() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}