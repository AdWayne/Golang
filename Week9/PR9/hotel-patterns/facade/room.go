package facade

import "fmt"

type RoomBookingSystem struct{}

func (r *RoomBookingSystem) BookRoom(room int) {
	fmt.Println("Бронирование номера:", room)
}

func (r *RoomBookingSystem) CancelRoom(room int) {
	fmt.Println("Отмена номера:", room)
}