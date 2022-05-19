package main

import (
	"fmt"
)

// Структура Утки
type Duck struct {
	flyer   Flyer
	quacker Quacker
}

type IDuck interface {
	Display()
	PerformFly()
	PerformQuack()
	Swim()
}

func (d *Duck) Display() {
	fmt.Println("Я утка")
}

func (d *Duck) PerformFly() {
	d.flyer.fly()
}

func (d *Duck) PerformQuack() {
	d.quacker.quack()
}

func (d *Duck) Swim() {
	fmt.Println("Все утки плавают")
}

// Интерфейс Flyer реализуется всеми классами, способными летать.
// От нового класса потребуется лишь реализация метода fly.
type Flyer interface {
	fly()
}

// Реализация fly для всех летающих уток
type FlyWithWings struct {
}

func (f FlyWithWings) fly() {
	fmt.Println("Я летаю")
}

// Реализация для всех уток не умеющих летать.
type FlyNoWay struct {
}

func (f FlyNoWay) fly() {
	fmt.Println("Я не летаю")
}
// Интерфейс Quacker реализуется всеми классами способными крякать.
// От нового класса потребуется лишь реализация метода quack
type Quacker interface {
	quack()
}
// Утки, которые крякают
type Quack struct {
}

func (q Quack) quack() {
	fmt.Println("Крякает")
}
// Утки, которые пищат 
type Squeak struct {
}

func (s Squeak) quack() {
	fmt.Println("Пищит")
}
// Утки, которые не издают ни звука
type MuteQuack struct {
}

func (m MuteQuack) quack() {
	fmt.Println("Не издает звуков")
}

type MallardDuck struct {
	Duck
}

func NewMallardDuck() IDuck {
	return &MallardDuck{Duck{flyer: FlyWithWings{}, quacker: Quack{}}}
}

func (m *MallardDuck) Display() {
	fmt.Println("Утка-кряква")
}

type RedHeadDuck struct {
	Duck
}

func NewRedHeadDuck() IDuck {
	return &RedHeadDuck{Duck{flyer: FlyWithWings{}, quacker: Quack{}}}
}

func (r *RedHeadDuck) Display() {
	fmt.Println("Рыжеголовая утка")
}

type RubberDuck struct {
	Duck
}

func NewRubberDuck() IDuck {
	return &RubberDuck{Duck{flyer: FlyNoWay{}, quacker: Squeak{}}}
}

func (r *RubberDuck) Display() {
	fmt.Println("Резиновая утка")
}

type DecoyDuck struct {
	Duck
}

func NewDecoyDuck() IDuck {
	return &DecoyDuck{Duck{flyer: FlyNoWay{}, quacker: MuteQuack{}}}
}

func (d *DecoyDuck) Display() {
	fmt.Println("Утка-приманка")
}

func main() {

	mallardDuck := NewMallardDuck()
	redHeadDuck := NewRedHeadDuck()
	rubberDuck := NewRubberDuck()
	decoyDuck := NewDecoyDuck()

	ducks := [4]IDuck{mallardDuck, redHeadDuck, rubberDuck, decoyDuck}

	for _, duck := range ducks {
		fmt.Printf("\n%T\n", duck)
		duck.Display()
		duck.PerformFly()
		duck.PerformQuack()
		duck.Swim()
	}
}
