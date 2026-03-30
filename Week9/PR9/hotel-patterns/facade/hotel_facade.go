package facade

import "fmt"

type HotelFacade struct {
	room      *RoomBookingSystem
	restaurant *RestaurantSystem
	event     *EventManagementSystem
	cleaning  *CleaningService
}

func NewHotelFacade() *HotelFacade {
	return &HotelFacade{
		room:      &RoomBookingSystem{},
		restaurant: &RestaurantSystem{},
		event:     &EventManagementSystem{},
		cleaning:  &CleaningService{},
	}
}

// Бронирование номера + еда + уборка
func (h *HotelFacade) BookRoomWithServices(room int) {
	fmt.Println("\n--- Бронирование номера с услугами ---")
	h.room.BookRoom(room)
	h.restaurant.OrderFood("Завтрак")
	h.cleaning.ScheduleCleaning(room)
}

// Организация мероприятия
func (h *HotelFacade) OrganizeEvent() {
	fmt.Println("\n--- Организация мероприятия ---")
	h.event.BookHall("Conference Hall")
	h.event.OrderEquipment("Projector")
	h.room.BookRoom(101)
	h.room.BookRoom(102)
}

// Ресторан + такси
func (h *HotelFacade) BookRestaurantWithTaxi() {
	fmt.Println("\n--- Ресторан + Такси ---")
	h.restaurant.BookTable(4)
	fmt.Println("Вызов такси...")
}