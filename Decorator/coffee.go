package main

import "fmt"

type beverager interface {
	Cost() float32
	GetDescription() string
}

type DarkRoast struct {
	description string
	cost        float32
}

func NewDarkRoast() beverager {
	dr := DarkRoast{description: "Dark Roast", cost: 0.99}
	return &dr
}

func (dr *DarkRoast) Cost() float32 {
	return dr.cost
}

func (dr *DarkRoast) GetDescription() string {
	return dr.description
}

type Decaf struct {
	description string
	cost        float32
}

func NewDecaf() beverager {
	d := Decaf{description: "Decaf", cost: 1.05}
	return &d
}

func (d *Decaf) Cost() float32 {
	return d.cost
}

func (d *Decaf) GetDescription() string {
	return d.description
}

type Espresso struct {
	description string
	cost        float32
}

func NewEspresso() beverager {
	e := Espresso{"Espresso", 1.99}
	return &e
}

func (e *Espresso) Cost() float32 {
	return e.cost
}

func (e *Espresso) GetDescription() string {
	return e.description
}

type HouseBlend struct {
	description string
	cost        float32
}

func NewHouseBlend() beverager {
	hb := HouseBlend{"House Blend", 0.89}
	return &hb
}

func (hb *HouseBlend) Cost() float32 {
	return hb.cost
}

func (hb *HouseBlend) GetDescription() string {
	return hb.description
}

type Mocha struct {
	beverage    beverager
	description string
	cost        float32
}

func NewMocha(b beverager) beverager {
	m := Mocha{b, "Mocha", 0.20}
	return &m
}

func (m *Mocha) Cost() float32 {
	return m.beverage.Cost() + m.cost
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", " + m.description
}

type Soy struct {
	beverage    beverager
	description string
	cost        float32
}

func NewSoy(b beverager) beverager {
	s := Soy{b, "Soy", 0.15}
	return &s
}

func (s *Soy) Cost() float32 {
	return s.beverage.Cost() + s.cost
}

func (s *Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", " + s.description
}

type SteamedMilk struct {
	beverage    beverager
	description string
	cost        float32
}

func NewSteamedMilk(b beverager) beverager {
	s := SteamedMilk{b, "Steamed Milk", 0.10}
	return &s 
}

func (s *SteamedMilk) Cost() float32 {
	return s.beverage.Cost() + s.cost 
}

func (s * SteamedMilk) GetDescription() string {
	return s.beverage.GetDescription() + ", " + s.description
}

type Whip struct {
	beverage beverager 
	description string 
	cost float32
}

func NewWhip(b beverager) beverager {
	w := Whip{b, "Whip", 0.10}
	return &w 
}

func (w *Whip) Cost() float32 {
	return w.beverage.Cost() + w.cost
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", " + w.description
}

func main() {
	es := NewEspresso()
	print(es)

	var b beverager
	b = NewDarkRoast()
	b = NewMocha(b)
	b = NewMocha(b)
	b = NewWhip(b)
	print(b)

	hbmw := NewWhip(
		NewMocha(
			NewSoy(
				NewHouseBlend(),
			),
		),
	)
	print(hbmw)
}

func print (b beverager) {
	fmt.Println(b.GetDescription(), "= \u20b9", b.Cost())
}