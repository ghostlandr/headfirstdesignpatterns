package pizzas

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Factory interface {
	CreatePizza(pizzaType string) Pizza
}
