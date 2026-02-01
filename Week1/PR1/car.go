package main

import "fmt"

type Car struct {
	vehicle      
	doors        int
	transmission string
}

func NewCar(make, model string, year, doors int, trans string) *Car {
	return &Car{
		vehicle:      vehicle{make, model, year},
		doors:        doors,
		transmission: trans,
	}
}

func (c *Car) GetDetails() string {
	return fmt.Sprintf("Машина: %s %s [%dг.], Дверей: %d, Трансмиссия: %s",
		c.make, c.model, c.year, c.doors, c.transmission)
}