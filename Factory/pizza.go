package main

import "fmt"

// Начинаем с абстрактного класса пиццы, все конкретные пиццы будут получены из этого класса.
type iPizza interface {
	prepare()
	bake()
	cut()
	box()
	getName() string
}

type pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

// Реализация стандартных этапов приготовления пиццы: приготовление, выпечка, нарезка, упаковка
func (p *pizza) prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.name)
	fmt.Printf("Tossing %s dough\n", p.dough)
	fmt.Printf("Adding %s sauce\n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for _, t := range p.toppings {
		fmt.Printf("\t%s, ", t)
	}
	fmt.Println()
}

func (p *pizza) bake() {
	fmt.Println("Baking Pizza for 25 min at 350")
}

func (p *pizza) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *pizza) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *pizza) getName() string {
	return p.name
}

type nyStyleCheesePizza struct {
	*pizza 
}

func newNYStyleCheesePizza() iPizza {
	p := &pizza {
		name: "NY Style Sauce and Cheese Pizza",
		dough: "Thin Crust Dough",
		sauce: "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}

	return nyStyleCheesePizza{
		pizza: p,
	}
}

type chicagoStyleCheesePizza struct {
	*pizza
}

func newChicagoStyleCheesePizza() iPizza {
	p := &pizza{
		name: "Chicago Style Deep Dish Cheese Pizza",
		dough: "Extra Thick Crust Dough",
		sauce: "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}

	return &chicagoStyleCheesePizza{
		pizza: p,
	}
}

// Нарезка в чикагском стиле предполагает нарезку квадратами
func (c *chicagoStyleCheesePizza) cut() {
	fmt.Println("Cutting the pizza into square slices")
}

// Абстрактный класс
type iPizzaStore interface {
	orderPizza(pizzaType string) iPizza
	createPizza(pizzaType string) (iPizza, error)
}

// Абстрактный конкретный тип
type aPizzaStore struct {
	createPizza func(pizzaType string) (iPizza, error)
}

func (a *aPizzaStore) orderPizza(pizzaType string) iPizza {
	if pizza, err := a.createPizza(pizzaType); err != nil {
		fmt.Println(err.Error())
		return nil 
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
		return pizza 
	}
}

// Каждый подкласс реализует метод createPizza() и использует метод orderPizza() в родительском классе.
type nyPizzaStore struct {
	*aPizzaStore
}

func newNYPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}

	nyPizzaStore := &nyPizzaStore{basePizzaStore}

	nyPizzaStore.aPizzaStore.createPizza = nyPizzaStore.createPizza

	return nyPizzaStore 
}

func (n *nyPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newNYStyleCheesePizza(), nil 
	}

	return nil, fmt.Errorf("Invalid pizza type")
}

type chicagoPizzaStore struct {
	*aPizzaStore
}

func newChicagoPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}

	chicagoPizzaStore := &chicagoPizzaStore{basePizzaStore}

	chicagoPizzaStore.aPizzaStore.createPizza = chicagoPizzaStore.createPizza

	return chicagoPizzaStore
}

func (c *chicagoPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newChicagoStyleCheesePizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}

func main() {
	nyPizzaStore := newNYPizzaStore()
	chicagoPizzaStore := newChicagoPizzaStore()

	pizza := nyPizzaStore.orderPizza("cheese")

	fmt.Printf("Ethan ordered %s pizza\n\n", pizza.getName())

	pizza = chicagoPizzaStore.orderPizza("cheese")

	fmt.Printf("Joel ordered %s pizza\n\n", pizza.getName())
}