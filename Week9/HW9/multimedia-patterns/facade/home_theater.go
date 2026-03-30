package facade

import "fmt"

type HomeTheaterFacade struct {
	tv      *TV
	audio   *AudioSystem
	dvd     *DVDPlayer
	console *GameConsole
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:      &TV{},
		audio:   &AudioSystem{},
		dvd:     &DVDPlayer{},
		console: &GameConsole{},
	}
}

// просмотр фильма
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("\n=== Просмотр фильма ===")

	h.tv.On()
	h.audio.On()

	h.audio.SetVolume(15)
	h.tv.SetChannel(1)

	h.dvd.Play(movie)
}

// выключить систему
func (h *HomeTheaterFacade) Shutdown() {
	fmt.Println("\n=== Выключение системы ===")

	h.dvd.Stop()
	h.audio.Off()
	h.tv.Off()
}

// запуск игры
func (h *HomeTheaterFacade) PlayGame(game string) {
	fmt.Println("\n=== Игровой режим ===")

	h.tv.On()
	h.console.On()
	h.console.PlayGame(game)
}

// режим музыки
func (h *HomeTheaterFacade) ListenMusic() {
	fmt.Println("\n=== Режим музыки ===")

	h.tv.On()
	h.audio.On()
	h.audio.SetVolume(20)
}

// регулировка громкости
func (h *HomeTheaterFacade) SetVolume(v int) {
	h.audio.SetVolume(v)
}