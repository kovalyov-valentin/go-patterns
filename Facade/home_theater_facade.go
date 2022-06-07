package main 

import (
	"fmt"
)

// Композиция:
// Это все компоненты подсистемы,
// которые будут использвоваться
type homeTheaterFacade struct {
	amp *amplifier
	dvd *dvdPlayer 
	proj *projector 
	lights *theaterLights 
	screen *screen 
	popper *popcornPopper
}

// Создание экземпляров всех компонентов подсистемы
func newHomeTheaterFacade() *homeTheaterFacade {
	return &homeTheaterFacade{
		amp: &amplifier{},
		dvd: &dvdPlayer{},
		proj: &projector{},
		lights: &theaterLights{},
		screen: &screen{},
		popper: &popcornPopper{},
	}
}

// Функция watchMovie() следует той же последовательности,
// что и раньше, но завершает ее удобным методом
// который выполняет всю работу.
// Следует обратить внимание, что для каждой задачи
// происходит делегирование ответственности
// соответствующему компоненты в подсистеме.
func (h *homeTheaterFacade) watchMovie(movie string) {
	fmt.Println("\nGet ready to watch a movie...")
	h.popper.on()
	h.popper.pop()
	h.lights.dim(10)
	h.screen.down()
	h.proj.on()
	h.proj.wideScreenMode()
	h.amp.on()
	h.amp.setSurroundSound()
	h.amp.setVolume(5)
	h.dvd.on()
	h.dvd.on()
	h.dvd.play(movie)
}

// endMovie() отвечает за завершение всех задач.
// Каждая задача делегируется соответствующему 
// компоненты в подсистеме.
func (h *homeTheaterFacade) endMovie() {
	fmt.Println("\nShutting movie theater down...")
	h.popper.off()
	h.lights.on()
	h.screen.up()
	h.proj.off()
	h.amp.off()
	h.dvd.stop()
	h.dvd.eject()
	h.dvd.off()
}