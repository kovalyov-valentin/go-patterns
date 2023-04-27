package main

import "fmt"

type pizza interface {
	getPrice() int
}

type peppyPaneer struct {
}

func (p *peppyPaneer) getPrice() int {
	return 20
}

type veggeMania struct {
}

func (v *veggeMania) getPrice() int {
	return 15
}

type cheeseToping struct {
	pizza pizza
}

func (c *cheeseToping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

type tomatoTopping struct {
	pizza pizza 
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice() 
	return pizzaPrice + 7
}

func main() {
	veggiePizza := &veggeMania{}

	veggiePizzaWithCheese := &cheeseToping{
		pizza: veggiePizza,
	}

	veggiePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.getPrice())

	peppyPaneerPizza := &peppyPaneer{} 

	peppyPaneerPizzaWithCheese := &cheeseToping{
		pizza: peppyPaneerPizza,
	}

	fmt.Printf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPaneerPizzaWithCheese.getPrice())
}