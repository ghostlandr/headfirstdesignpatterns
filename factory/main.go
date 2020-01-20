// Implementing the pattern from chapter 4 (Factory)
package main

import (
	"headfirstdesigntraining/factory/pizzastore"
)

func main() {
	chiPiStore := pizzastore.ChicagoStylePizzaStore()

	chiPiStore.OrderPizza("cheese")

	nyPiStore := pizzastore.NewYorkStylePizzaStore()

	nyPiStore.OrderPizza("pepperoni")
}
