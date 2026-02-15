package main

import "fmt"



type Motorcycle struct {
	vehicle
	bodyType string
	hasBox   bool
}

func NewMotorcycle(make, model string, year int, body string, box bool) *Motorcycle {
	return &Motorcycle{
		vehicle:  vehicle{make, model, year},
		bodyType: body,
		hasBox:   box,
	}
}

func (m *Motorcycle) GetDetails() string {
	boxStatus := "без багажника"
	if m.hasBox {
		boxStatus = "с багажником"
	}
	return fmt.Sprintf("Мотоцикл: %s %s [%dг.], Тип: %s, %s",
		m.make, m.model, m.year, m.bodyType, boxStatus)
}