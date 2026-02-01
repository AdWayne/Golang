package main

import "fmt"

func main() {
	// Инициализация объектов
	porsche := NewCar("Porsche", "911", 2023, 2, "PDK")
	ducati := NewMotorcycle("Ducati", "V4", 2024, "Superbike", false)
	toyota := NewCar("Toyota", "Camry", 2022, 4, "Automatic")
	// Работа с гаражом
	myGarage := &Garage{Name: "VIP-Box"}
	myGarage.Add(porsche)
	myGarage.Add(ducati)
	myGarage.Add(toyota)

	// Работа с автопарком
	fleet := &Fleet{}
	fleet.AddGarage(myGarage)

	// Тестирование полиморфизма
	fmt.Println("Список транспорта:")
	for _, v := range myGarage.Vehicles {
		fmt.Println(v.GetDetails())
	}

	// Поиск
	fleet.Search("911")
	fleet.DeleteCar("Camry")
	fleet.DeleteGarage("VIP-Box")
}