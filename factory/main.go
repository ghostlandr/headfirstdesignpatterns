// Implementing the pattern from chapter 4 (Factory)
package main

import (
	"headfirstdesigntraining/factory/pizzastore"
)

func main() {
	piStore := pizzastore.ChicagoStylePizzaStore()

	piStore.OrderPizza("cheese")
	piStore.OrderPizza("pepperoni")

	piStore = pizzastore.NewYorkStylePizzaStore()

	piStore.OrderPizza("pepperoni")
	piStore.OrderPizza("cheese")
}
