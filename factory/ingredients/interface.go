package ingredients

type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateVeggies() []Veggies
	CreatePepperoni() Pepperoni
	CreateClam() Clam
}

type Dough interface {
	Name() string
}

type Sauce interface {
	Name() string
}

type Cheese interface {
	Name() string
}

type Veggies interface {
	Name() string
}

type Pepperoni interface {
	Name() string
}

type Clam interface {
	Name() string
}
