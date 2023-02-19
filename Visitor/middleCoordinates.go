package main 

import (
	"fmt"
)

type middleCoordinates struct {
	x int 
	y int 
}

func (c *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (c *middleCoordinates) visitForCircle(s *circle) {
	fmt.Println("Calculating middle point coordinates gor circle")
}

func (c *middleCoordinates) visitForRestangle(s *restangle) {
	fmt.Println("Calculatinf middle point coordinates for restangle")
}
