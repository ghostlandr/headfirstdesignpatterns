package ingredients

func NYIngredientFactory() PizzaIngredientFactory {
	return nyIngredientFactory{}
}

type nyIngredientFactory struct {}

func (f nyIngredientFactory) CreateDough() Dough {
	return thickCrustDough{name: "Thin dough"}
}

func (f nyIngredientFactory) CreateSauce() Sauce {
	return plumTomatoSauce{name: "marinara"}
}

func (f nyIngredientFactory) CreateCheese() Cheese {
	return mozzarellaCheese{name: "cheese"}
}

func (f nyIngredientFactory) CreateVeggies() []Veggies {
	return []Veggies{}
}

func (f nyIngredientFactory) CreatePepperoni() Pepperoni {
	return slicedPepperoni{name: "sliced pepperoni"}
}

func (f nyIngredientFactory) CreateClam() Clam {
	return frozenClam{name: "fresh clams"}
}

