package main

import (
	"fmt"
)

type duck interface {
	quack()
	fly()
}

type mallarDuck struct {}

func (m *mallarDuck) quack() {
	fmt.Println("Quack")
}

func (m *mallarDuck) fly() {
	fmt.Println("I'm flying")
}

type turkey interface {
	gobble()
	fly()
}

type wildTurkey struct{}

func (w *wildTurkey) gobble() {
	fmt.Println("Gobble Gobble")
}

func (w *wildTurkey) fly() {
	fmt.Println("I'm flying a short distance")
}

type turkeyAdapter struct {
	turkey turkey
}

func newTurkeyAdapter(t turkey) *turkeyAdapter {
	return &turkeyAdapter{
		turkey: t,
	}
}

func (t *turkeyAdapter) quack() {
	t.turkey.gobble()
}

func (t *turkeyAdapter) fly() {
	for i := 1; i <=5; i++ {
		t.turkey.fly()
	}
}

func main() {
	// Создание утки и индейки
	mallarDuck := &mallarDuck{}

	wildTurkey := &wildTurkey{}

	// Заворачиваем индейку в адаптер для индейки,
	// что делает ее похожей на утку
	turkeyAdapter := newTurkeyAdapter(wildTurkey)

	//Передача duck методу testDuck(), который ожидает объект Duck
	fmt.Println("The Duck says...")
	testDuck(mallarDuck)

	// Выдать индейку за утку тем же способом
	fmt.Println("\nThe Turkey Adapter says...")
	testDuck(turkeyAdapter)
}

func testDuck(d duck) {
	d.quack()
	d.fly()
}
