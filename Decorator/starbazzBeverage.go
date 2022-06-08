package main

import "fmt"

// В нашем случае мы можем использовать интерфейс для напитков вместо абстрактного класса.
type beverage interface {
	description() string
	cost() float32
}

type espresso struct {
}

func (e *espresso) description() string {
	return "Espresso"
}

func (e *espresso) cost() float32 {
	return 1.99
}

type houseBlend struct {
}

func (hb *houseBlend) description() string {
	return "House Blend Coffee"
}

func (hb *houseBlend) cost() float32 {
	return 0.99
}

type darkRoast struct {
}

func (dr *darkRoast) description() string {
	return "Dark Roast Coffee"
}

func (dr *darkRoast) cost() float32 {
	return 0.99
}

type decaf struct {
}

func (d *decaf) description() string {
	return "Decaf"
}

func (d *decaf) cost() float32 {
	return 1.05
}

type mocha struct {
	beverage beverage
}

func (m *mocha) description() string {
	return m.beverage.description() + ", Mocha"
}

func (m *mocha) cost() float32 {
	return m.beverage.cost() + .2
}

type milk struct {
	beverage beverage
}

func (m *milk) description() string {
	return m.beverage.description() + ", Milk"
}

func (m *milk) cost() float32 {
	return m.beverage.cost() + .1 
}

type soy struct {
	beverage beverage
}

func (s *soy) description() string {
	return s.beverage.description() + ", Soy"
}

func (s *soy) cost() float32 {
	return s.beverage.cost() + .15
}

type whip struct {
	beverage beverage
}

func (w *whip) description() string {
	return w.beverage.description() + ", Whip"
}

func (w *whip) cost() float32 {
	return w.beverage.cost() + .1
}

func main() {
	beverage := &espresso{}

	fmt.Printf("%s $%.2f\n", beverage.description(), beverage.cost())

	darkRoast := &darkRoast{}

	singleMocha := &mocha {
		beverage: darkRoast,
	}

	doubleMocha := &mocha {
		beverage: singleMocha,
	}

	doubleMochaWhip := &whip {
		beverage: doubleMocha,
	}

	fmt.Printf("%s $%.2f\n", doubleMochaWhip.description(), doubleMochaWhip.cost())

	soyMochaWhipHouseBlend := &whip {
		beverage: &mocha {
			beverage: &soy {
				beverage: &houseBlend {},
			},
		},
	}

	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.description(), soyMochaWhipHouseBlend.cost())
}