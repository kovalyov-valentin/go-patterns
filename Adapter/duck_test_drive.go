package main

import (
	"fmt"
)

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
