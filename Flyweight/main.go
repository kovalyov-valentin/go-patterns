package main 

import (
	"fmt"
)

func main() {
	// game := newGame()
	// //Add Terrorist 
	// game.addTerrorist(TerroristDressType)
	// game.addTerrorist(TerroristDressType)
	// game.addTerrorist(TerroristDressType)
	// game.addTerrorist(TerroristDressType)
	// //Add CounterTerrorist
	// game.addCounterTerrorist(CounterTerroristDressType)
	// game.addCounterTerrorist(CounterTerroristDressType)
	// game.addCounterTerrorist(CounterTerroristDressType)
	terrorist := newTerroristDress()
	fmt.Println(terrorist)
	counterTerrorist := newCounterTerroristDress()
	fmt.Println(counterTerrorist)
	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}