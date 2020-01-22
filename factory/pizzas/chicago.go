package pizzas

import (
	"fmt"

	"headfirstdesigntraining/factory/ingredients"
)

func ChicagoStyleFactory() Factory {
	return chicagoFactory{}
}

type chicagoFactory struct {}

func (c chicagoFactory) CreatePizza(pizzaType string) Pizza {
	fmt.Printf("Preparing a delectable Chicago-style %s pizza\n", pizzaType)
	f := ingredients.ChicagoIngredientFactory()
	switch pizzaType {
	case "cheese":
		return Cheese("Chicago Cheese", f)
	case "pepperoni":
		return Pepperoni("Chicago Pepperoni", f)
	}
	return nil
}
