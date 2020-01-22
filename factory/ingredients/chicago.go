package ingredients

func ChicagoIngredientFactory() PizzaIngredientFactory {
	return chicagoIngredientFactory{}
}

type chicagoIngredientFactory struct {}

func (c chicagoIngredientFactory) CreateDough() Dough {
	return thickCrustDough{name: "Thicc dough"}
}

func (c chicagoIngredientFactory) CreateSauce() Sauce {
	return plumTomatoSauce{name: "plum tomato sauce"}
}

func (c chicagoIngredientFactory) CreateCheese() Cheese {
	return mozzarellaCheese{name: "cheese"}
}

func (c chicagoIngredientFactory) CreateVeggies() []Veggies {
	return []Veggies{eggplant{name: "eggplant"}, blackOlives{name: "black olives"}}
}

func (c chicagoIngredientFactory) CreatePepperoni() Pepperoni {
	return slicedPepperoni{name: "sliced pepperoni"}
}

func (c chicagoIngredientFactory) CreateClam() Clam {
	return frozenClam{name: "unfortunately frozen clams"}
}

