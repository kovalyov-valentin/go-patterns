package main 

import (
	"fmt"
)

func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}
	restangle := &restangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	restangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	restangle.accept(middleCoordinates) 
}