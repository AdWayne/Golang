package exchange

import "fmt"

type MobileApp struct {
	Name string
}

func (m *MobileApp) Update(rate float64) {
	fmt.Printf("📱 %s получил обновление курса: %.2f\n", m.Name, rate)
}