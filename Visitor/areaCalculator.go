package main 

import (
	"fmt"
)

type areaCalculator struct {
	area int 
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForRestangle(s *restangle) {
	fmt.Println("Calculating area for restangle")
}