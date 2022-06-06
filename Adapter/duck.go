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
