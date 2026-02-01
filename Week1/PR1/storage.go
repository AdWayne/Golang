package main

import (
	"fmt"
	"strings"
)

type Garage struct {
	Name     string
	Vehicles []VehicleInterface
}

func (g *Garage) Add(v VehicleInterface) {
	g.Vehicles = append(g.Vehicles, v)
}

// Fleet управляет списком гаражей
type Fleet struct {
	Garages []*Garage
}

func (f *Fleet) AddGarage(g *Garage) {
	f.Garages = append(f.Garages, g)
}

func (f *Fleet) Search(modelName string) {
	fmt.Printf("\n--- Результаты поиска для '%s' ---\n", modelName)
	for _, g := range f.Garages {
		for _, v := range g.Vehicles {
			if strings.EqualFold(v.GetModel(), modelName) {
				fmt.Printf("Найдено в гараже [%s]: %s\n", g.Name, v.GetDetails())
			}
		}
	}
}

func (f *Fleet) DeleteGarage(name string) {
	initialLen := len(f.Garages)
	newGarages := []*Garage{}
	
	for _, g := range f.Garages {
		if g.Name != name {
			newGarages = append(newGarages, g)
		}
	}
	
	f.Garages = newGarages

	if len(f.Garages) < initialLen {
		fmt.Printf("-> Гараж [%s] успешно удален из автопарка.\n", name)
	} else {
		fmt.Printf("-> Ошибка: Гараж [%s] не найден.\n", name)
	}
}

func (g *Garage) RemoveVehicleByModel(modelName string) bool {
	initialLen := len(g.Vehicles)
	newVehicles := []VehicleInterface{}

	for _, v := range g.Vehicles {
		// Используем метод GetModel из нашего интерфейса
		if v.GetModel() != modelName {
			newVehicles = append(newVehicles, v)
		}
	}

	g.Vehicles = newVehicles
	return len(g.Vehicles) < initialLen
}

func (f *Fleet) DeleteCar(modelName string) {
	fmt.Printf("\nПопытка удаления транспорта '%s' из всех гаражей...\n", modelName)
	foundTotal := false

	for _, g := range f.Garages {
		if g.RemoveVehicleByModel(modelName) {
			fmt.Printf("  [OK] Машина '%s' удалена из гаража [%s]\n", modelName, g.Name)
			foundTotal = true
		}
	}

	if !foundTotal {
		fmt.Printf("  [!] Машина '%s' не найдена ни в одном гараже.\n", modelName)
	}
}