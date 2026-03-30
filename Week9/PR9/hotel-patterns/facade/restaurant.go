package facade

import "fmt"

type RestaurantSystem struct{}

func (r *RestaurantSystem) BookTable(count int) {
	fmt.Println("Бронирование стола на", count, "человек")
}

func (r *RestaurantSystem) OrderFood(food string) {
	fmt.Println("Заказ еды:", food)
}