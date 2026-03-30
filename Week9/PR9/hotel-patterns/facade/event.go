package facade

import "fmt"

type EventManagementSystem struct{}

func (e *EventManagementSystem) BookHall(name string) {
	fmt.Println("Бронирование зала:", name)
}

func (e *EventManagementSystem) OrderEquipment(eq string) {
	fmt.Println("Заказ оборудования:", eq)
}