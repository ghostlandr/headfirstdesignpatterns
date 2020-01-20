package pizzastore

import (
	"headfirstdesigntraining/factory/pizzas"
)

type PizzaStore interface {
	OrderPizza(pizzaType string) pizzas.Pizza
	CreatePizza(pizzaType string) pizzas.Pizza
}
