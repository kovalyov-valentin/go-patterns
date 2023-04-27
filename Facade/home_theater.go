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

type amplifier struct{}

func (a *amplifier) on() {
	fmt.Println("Top-O-Line Amplifier on")
}

func (a *amplifier) setSurroundSound() {
	fmt.Println("Top-O-Line Amplifier surround sound on (5 speakers, 1 subwoofer)")
}

func (a *amplifier) setVolume(n int) {
	fmt.Printf("Top-O-Line Amplifier setting volume to %d\n", n)
}

func (a *amplifier) off() {
	fmt.Println("Top-O-Line Amplifier off")
}

type dvdPlayer struct {}

func (d *dvdPlayer) on() {
	fmt.Println("Top-O-Line DVD Player on")
}

func (d *dvdPlayer) play(movie string) {
	fmt.Printf("Top-O-Line DVD Player playing \"%s\"\n", movie)
}

func (d *dvdPlayer) stop() {
	fmt.Println("Top-O-Line DVD Player stopped")
}

func (d *dvdPlayer) eject() {
	fmt.Println("Top-O-Line DVD Player eject")
}

func (d *dvdPlayer) off() {
	fmt.Println("Top-O-Line DVD Player off")
}

type projector struct{}

func (p *projector) on() {
	fmt.Println("Top-O-Line Projector on")
}

func (p *projector) wideScreenMode() {
	fmt.Println("Top-O-Line Projector in widescreen mode (16x9 aspect ratio)")
}

func (p *projector) off() {
	fmt.Println("Top-O-Line Pojector off")
}

type theaterLights struct{}

func (t *theaterLights) dim(n int) {
	fmt.Printf("Theater Ceiling Lights dimming to %d%%\n", n)
}

func (t *theaterLights) on() {
	fmt.Println("Theater Ceiling Lights on")
}

type screen struct {}

func (s *screen) down() {
	fmt.Println("Theater screen going down")
}

func (s *screen) up() {
	fmt.Println("Theater screen going up")
}

type popcornPopper struct {}

func (p *popcornPopper) on() {
	fmt.Println("Popcorn Popper popping popcorn!")
}

func (p *popcornPopper) pop() {
	fmt.Println("Popcorn Popper popping popcorn!")
}

func (p *popcornPopper) off() {
	fmt.Println("Popcorn Popper off")
}

func main() {
	homeTheater := newHomeTheaterFacade()

	homeTheater.watchMovie("Raiders of the Lost Ark")
	homeTheater.endMovie()
}
