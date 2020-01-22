package pizzas

import (
	"fmt"

	"headfirstdesigntraining/factory/ingredients"
)

func NewYorkStyleFactory() Factory {
	return nyFactory{}
}

type nyFactory struct {}

func (c nyFactory) CreatePizza(pizzaType string) Pizza {
	fmt.Printf("Preparing a classy NY-style %s pizza\n", pizzaType)
	f := ingredients.NYIngredientFactory()
	switch pizzaType {
	case "cheese":
		return nyCheese(f)
	case "pepperoni":
		return nyPepperoni{}
	}
	return nil
}

func nyCheese(f ingredients.PizzaIngredientFactory) Pizza {
	return Cheese("New York Cheese", f)
}

type nyPepperoni struct{}

func (p nyPepperoni) Prepare() {
	fmt.Print("Slapping some dough...\nSpreading some sauce\nSprinkling some cheese\nPlacing some pepps\n")
}

func (p nyPepperoni) Bake() {
	fmt.Print("Sliding this baby into the oven\n")
}

func (p nyPepperoni) Cut() {
	fmt.Print("Cutting this sucker into quarters for your folding pleasure\n")
}

func (p nyPepperoni) Box() {
	fmt.Print("Boxing 'er up\n")
}
