package main

import "fmt"

// VehicleInterface — контракт для всех ТС (Полиморфизм)
type VehicleInterface interface {
	Start()
	Stop()
	GetDetails() string
	GetModel() string
}

// Базовая структура (скрытые поля = Инкапсуляция)
type vehicle struct {
	make  string
	model string
	year  int
}

func (v vehicle) Start() {
	fmt.Printf("Система: %s %s заведен.\n", v.make, v.model)
}

func (v vehicle) Stop() {
	fmt.Printf("Система: %s %s заглушен.\n", v.make, v.model)
}

func (v vehicle) GetModel() string {
	return v.model
}