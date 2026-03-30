package facade

import "fmt"

type CleaningService struct{}

func (c *CleaningService) ScheduleCleaning(room int) {
	fmt.Println("Уборка номера:", room)
}