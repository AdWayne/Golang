package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ===== 1) Интерфейс продукта =====
type IVehicle interface {
	Drive()
	Refuel()
	Info() string
}

// ===== 2) Конкретные продукты =====

// Car: марка, модель, тип топлива
type Car struct {
	brand    string
	model    string
	fuelType string
}

func (c Car) Drive()  { fmt.Printf("🚗 %s едет: %s %s на %s.\n", c.Info(), c.brand, c.model, c.fuelType) }
func (c Car) Refuel() { fmt.Printf("⛽ %s заправляется: топливо = %s.\n", c.Info(), c.fuelType) }
func (c Car) Info() string {
	return fmt.Sprintf("Car(%s %s, fuel=%s)", c.brand, c.model, c.fuelType)
}

// Motorcycle: тип (sport/touring), объем двигателя
type Motorcycle struct {
	motoType   string
	engineCC   int
}

func (m Motorcycle) Drive()  { fmt.Printf("🏍️ %s едет: тип=%s, %dcc.\n", m.Info(), m.motoType, m.engineCC) }
func (m Motorcycle) Refuel() { fmt.Printf("⛽ %s заправляется.\n", m.Info()) }
func (m Motorcycle) Info() string {
	return fmt.Sprintf("Motorcycle(type=%s, engine=%dcc)", m.motoType, m.engineCC)
}

// Truck: грузоподъемность, количество осей
type Truck struct {
	payloadKg int
	axles     int
}

func (t Truck) Drive()  { fmt.Printf("🚚 %s едет: грузоподъемность=%dкг, осей=%d.\n", t.Info(), t.payloadKg, t.axles) }
func (t Truck) Refuel() { fmt.Printf("⛽ %s заправляется (дизель условно).\n", t.Info()) }
func (t Truck) Info() string {
	return fmt.Sprintf("Truck(payload=%dkg, axles=%d)", t.payloadKg, t.axles)
}

// ===== 7) Расширение: Bus =====
// Bus: вместимость, маршрут
type Bus struct {
	capacity int
	route    string
}

func (b Bus) Drive()  { fmt.Printf("🚌 %s едет по маршруту %q.\n", b.Info(), b.route) }
func (b Bus) Refuel() { fmt.Printf("⛽ %s заправляется.\n", b.Info()) }
func (b Bus) Info() string {
	return fmt.Sprintf("Bus(capacity=%d, route=%s)", b.capacity, b.route)
}

// ===== 3) Абстрактная фабрика (Factory Method) =====
type VehicleFactory interface {
	CreateVehicle() IVehicle
}

// ===== 4) Конкретные фабрики (с параметрами) =====

type CarFactory struct {
	brand    string
	model    string
	fuelType string
}
func (f CarFactory) CreateVehicle() IVehicle {
	return Car{brand: f.brand, model: f.model, fuelType: f.fuelType}
}

type MotorcycleFactory struct {
	motoType string
	engineCC int
}
func (f MotorcycleFactory) CreateVehicle() IVehicle {
	return Motorcycle{motoType: f.motoType, engineCC: f.engineCC}
}

type TruckFactory struct {
	payloadKg int
	axles     int
}
func (f TruckFactory) CreateVehicle() IVehicle {
	return Truck{payloadKg: f.payloadKg, axles: f.axles}
}

type BusFactory struct {
	capacity int
	route    string
}
func (f BusFactory) CreateVehicle() IVehicle {
	return Bus{capacity: f.capacity, route: f.route}
}

// ===== CLI helpers =====

type App struct {
	in       *bufio.Reader
	vehicles []IVehicle
}

func NewApp() *App {
	return &App{in: bufio.NewReader(os.Stdin)}
}

func (a *App) readLine(prompt string) string {
	for {
		fmt.Print(prompt)
		s, _ := a.in.ReadString('\n')
		s = strings.TrimSpace(s)
		if s != "" {
			return s
		}
		fmt.Println("Введите непустое значение.")
	}
}

func (a *App) readInt(prompt string, min int, max int) int {
	for {
		s := a.readLine(prompt)
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Нужно целое число.")
			continue
		}
		if v < min || v > max {
			fmt.Printf("Число должно быть в диапазоне [%d..%d].\n", min, max)
			continue
		}
		return v
	}
}

func (a *App) createVehicleByUserInput() {
	fmt.Println("\nВыберите тип транспорта:")
	fmt.Println("1) Car")
	fmt.Println("2) Motorcycle")
	fmt.Println("3) Truck")
	fmt.Println("4) Bus (новый тип)")

	choice := a.readInt("Ваш выбор (1-4): ", 1, 4)

	var factory VehicleFactory

	switch choice {
	case 1:
		brand := a.readLine("Марка: ")
		model := a.readLine("Модель: ")
		fuel := a.readLine("Тип топлива (бензин/дизель/электро/газ): ")
		factory = CarFactory{brand: brand, model: model, fuelType: fuel}

	case 2:
		motoType := a.readLine("Тип мотоцикла (sport/touring/other): ")
		engine := a.readInt("Объем двигателя (cc, например 600): ", 50, 5000)
		factory = MotorcycleFactory{motoType: motoType, engineCC: engine}

	case 3:
		payload := a.readInt("Грузоподъемность (кг): ", 1, 200000)
		axles := a.readInt("Количество осей: ", 2, 12)
		factory = TruckFactory{payloadKg: payload, axles: axles}

	case 4:
		cap := a.readInt("Вместимость (чел): ", 5, 200)
		route := a.readLine("Маршрут (например: 12A / Центр—Вокзал): ")
		factory = BusFactory{capacity: cap, route: route}
	}

	v := factory.CreateVehicle()
	a.vehicles = append(a.vehicles, v)
	fmt.Println("✅ Создано:", v.Info())
}

func (a *App) listVehicles() {
	if len(a.vehicles) == 0 {
		fmt.Println("\nСписок пуст.")
		return
	}
	fmt.Println("\nВаш транспорт:")
	for i, v := range a.vehicles {
		fmt.Printf("%d) %s\n", i+1, v.Info())
	}
}

func (a *App) actionOnVehicle() {
	if len(a.vehicles) == 0 {
		fmt.Println("\nСначала создайте транспорт.")
		return
	}
	a.listVehicles()
	idx := a.readInt("Выберите номер транспорта: ", 1, len(a.vehicles)) - 1

	fmt.Println("\nДействие:")
	fmt.Println("1) Drive()")
	fmt.Println("2) Refuel()")
	act := a.readInt("Ваш выбор (1-2): ", 1, 2)

	switch act {
	case 1:
		a.vehicles[idx].Drive()
	case 2:
		a.vehicles[idx].Refuel()
	}
}

func (a *App) Run() {
	for {
		fmt.Println("\n==== Transport System (Factory Method) ====")
		fmt.Println("1) Создать транспорт")
		fmt.Println("2) Показать список")
		fmt.Println("3) Выполнить действие (Drive/Refuel)")
		fmt.Println("4) Выход")

		cmd := a.readInt("Команда (1-4): ", 1, 4)

		switch cmd {
		case 1:
			a.createVehicleByUserInput()
		case 2:
			a.listVehicles()
		case 3:
			a.actionOnVehicle()
		case 4:
			fmt.Println("Пока!")
			return
		}
	}
}

func main() {
	app := NewApp()
	app.Run()
}
