package pizzas

import (
	"fmt"
)

type chicagoFactory struct {}

func (c chicagoFactory) CreatePizza(pizzaType string) Pizza {
	fmt.Printf("Preparing a delectable Chicago-style %s pizza\n", pizzaType)
	switch pizzaType {
	case "cheese":
		return chicagoCheese{}
	case "pepperoni":
		return chicagoPepperoni{}
	}
	return nil
}

type chicagoCheese struct{}

func (p chicagoCheese) Prepare() {
	fmt.Print("Pounding out a monster dough ball\nDumping out some sauce\nPouring a mountain of cheese on top\n")
}

func (p chicagoCheese) Bake() {
	fmt.Print("Finding two other people to help me get this into the oven\n")
}

func (p chicagoCheese) Cut() {
	fmt.Print("Firing up the band saw to cut this thing\n")
}

func (p chicagoCheese) Box() {
	fmt.Print("Using our new fridge box to be able to hold this behemoth\n")
}

type chicagoPepperoni struct{}

func (p chicagoPepperoni) Prepare() {
}

func (p chicagoPepperoni) Bake() {
}

func (p chicagoPepperoni) Cut() {
}

func (p chicagoPepperoni) Box() {
}
