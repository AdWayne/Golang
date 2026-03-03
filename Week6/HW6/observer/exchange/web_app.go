package exchange

import "fmt"

type WebApp struct {
	Site string
}

func (w *WebApp) Update(rate float64) {
	fmt.Printf("🌐 Сайт %s обновил курс: %.2f\n", w.Site, rate)
}