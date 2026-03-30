package main

import (
	"fmt"
	"hotel-patterns/composite"
	"hotel-patterns/facade"
)

func main() {

	// ================= FACADE =================
	hotel := facade.NewHotelFacade()

	hotel.BookRoomWithServices(101)
	hotel.OrganizeEvent()
	hotel.BookRestaurantWithTaxi()

	// ================= COMPOSITE =================

	fmt.Println("\n===== ORGANIZATION STRUCTURE =====")

	dev := composite.NewDepartment("Development")

	dev.Add(&composite.Employee{
		Name:     "Ali",
		Position: "Backend",
		Salary:   2000,
	})

	dev.Add(&composite.Employee{
		Name:     "Dana",
		Position: "Frontend",
		Salary:   1800,
	})

	dev.Add(&composite.Contractor{
		Name: "Freelancer",
		Pay:  500,
	})

	company := composite.NewDepartment("Company")
	company.Add(dev)

	company.Show("")

	fmt.Println("\nTotal salary:", company.GetSalary())
	fmt.Println("Total employees:", company.GetEmployeeCount())
}