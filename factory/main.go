// Implementing the pattern from chapter 4 (Factory)
package main

import (
	"headfirstdesigntraining/factory/pizzas"
	"headfirstdesigntraining/factory/pizzastore"
)

func main() {
	piStore := pizzastore.ChicagoStylePizzaStore()

	piStore.OrderPizza("cheese")
	piStore.OrderPizza("pepperoni")

	piStore = pizzastore.NewYorkStylePizzaStore()

	piStore.OrderPizza("pepperoni")
	piStore.OrderPizza("cheese")

	piStore = pizzastore.NewPizzaStore(pizzas.ChicagoStyleFactory())

	piStore.OrderPizza("cheese")
	piStore.OrderPizza("pepperoni")
}
