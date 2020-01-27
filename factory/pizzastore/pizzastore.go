package pizzastore

import (
	"headfirstdesigntraining/factory/pizzas"
)

func NewPizzaStore(f pizzas.Factory) PizzaStore {
	return &pizzaStore{
		Factory: f,
	}
}

type pizzaStore struct{
	pizzas.Factory
}

func (p pizzaStore) OrderPizza(pizzaType string) pizzas.Pizza {
	pizza := p.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

func ChicagoStylePizzaStore() PizzaStore {
	return chicagoStylePizzaStore{
		pizzas.ChicagoStyleFactory(),
	}
}

type chicagoStylePizzaStore struct{
	pizzas.Factory
}

func (p chicagoStylePizzaStore) OrderPizza(pizzaType string) pizzas.Pizza {
	pizza := p.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

func NewYorkStylePizzaStore() PizzaStore {
	return nyStylePizzaStore{
		pizzas.NewYorkStyleFactory(),
	}
}

type nyStylePizzaStore struct{
	pizzas.Factory
}

func (p nyStylePizzaStore) OrderPizza(pizzaType string) pizzas.Pizza {
	pizza := p.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
