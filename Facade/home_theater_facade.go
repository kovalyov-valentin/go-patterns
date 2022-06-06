package main 

import (
	"fmt"
)

type homeTheaterFacade struct {
	amp *amplifier
	dvd *dvdPlayer 
	proj *projector 
	lights *theaterLights 
	screen *screen 
	popper *popcornPopper
}

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