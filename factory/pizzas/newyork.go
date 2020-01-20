package pizzas

import (
	"fmt"
)

type nyFactory struct {}

func (c nyFactory) CreatePizza(pizzaType string) Pizza {
	fmt.Printf("Preparing a classy NY-style %s pizza\n", pizzaType)
	switch pizzaType {
	case "cheese":
		return nyCheese{}
	case "pepperoni":
		return nyPepperoni{}
	}
	return nil
}

type nyCheese struct{}

func (p nyCheese) Prepare() {
}

func (p nyCheese) Bake() {
}

func (p nyCheese) Cut() {
}

func (p nyCheese) Box() {
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
	fmt.Print("Boxing 'er up")
}
