package factory_method

import (
	"log"
	"testing"
)

type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

type Creator interface {
	CreateProduct(action action) Product
}

type Product interface {
	Use() string
}

type ConcreteCreator struct {
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product 

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default: 
		log.Fatalln("Unknown ACtion")
	}

	return product 
}

type ConcreteProductA struct {
	action string 
}

func (p *ConcreteProductA) Use() string {
	return p.action
}

type ConcreteProductB struct {
	action string 
}

func (p *ConcreteProductB) Use() string {
	return p.action 
}

type ConcreteProductC struct {
	action string 
}

func (p *ConcreteProductC) Use() string {
	return p.action 
}

func TestFactoryMethod(t *testing.T) {
	assert := []string{"A", "B", "C"}

	factory := NewCreator()
	products := []Product{
		factory.CreateProduct(A), 
		factory.CreateProduct(B),
		factory.CreateProduct(C),
	}

	for i, product := range products {
		if action := product.Use(); action != assert[i] {
			t.Errorf("Expect action to %s, but %s.\n", assert[i], action)
		}
	}
}