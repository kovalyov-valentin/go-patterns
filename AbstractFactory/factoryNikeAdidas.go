package main

import (
	"fmt"
)

type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 14,
		},
	}
}

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 14,
		},
	}
}

type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string 
	getSize() int
}

type shoe struct {
	logo string 
	size int 
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo 
}

func (s *shoe) getLogo() string {
	return s.logo 
}

func (s *shoe) setSize(size int) {
	s.size = size 
}

func (s *shoe) getSize() int {
	return s.size 
}

type adidasShoe struct {
	shoe 
}

type nikeShoe struct {
	shoe 
}

type iShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string 
	getSize() int
}

type short struct {
	logo string 
	size int 
}

func (s *short) setLogo(logo string) {
	s.logo = logo 
}

func (s *short) getLogo() string {
	return s.logo 
}

func (s *short) setSize(size int) {
	s.size = size 
}

func (s *short) getSize() int {
	return s.size 
}

type adidasShort struct {
	short 
}

type nikeShort struct {
	short 
}

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()

	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}