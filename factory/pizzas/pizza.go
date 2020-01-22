package pizzas

import (
	"fmt"

	"headfirstdesigntraining/factory/ingredients"
)

func Pepperoni(name string, f ingredients.PizzaIngredientFactory) Pizza{
	return &pepperoni{ingredientFactory: f, name: name}
}

type pepperoni struct{
	name string
	ingredientFactory ingredients.PizzaIngredientFactory
}

func (p pepperoni) Prepare() {
	fmt.Println("Preparing", p.name)
	// Use the ingredients from the ingredient factory instead. Now we only need one pepperoni pizza
	// struct for all the types of pepperoni
	d := p.ingredientFactory.CreateDough()
	s := p.ingredientFactory.CreateSauce()
	c := p.ingredientFactory.CreateCheese()
	pep := p.ingredientFactory.CreatePepperoni()
	fmt.Printf("Laying on the ingredients: %s, %s, %s, %s\n", d.Name(), s.Name(), c.Name(), pep.Name())
}

func (p pepperoni) Bake() {
}

func (p pepperoni) Cut() {
}

func (p pepperoni) Box() {
}

func Cheese(name string, f ingredients.PizzaIngredientFactory) Pizza{
	return &cheese{ingredientFactory: f, name: name}
}

type cheese struct{
	name string
	ingredientFactory ingredients.PizzaIngredientFactory
}

func (c cheese) Prepare() {
	fmt.Println("Preparing", c.name)
	// Use the ingredients from the ingredient factory instead. Now we only need one pepperoni pizza
	// struct for all the types of pepperoni
	d := c.ingredientFactory.CreateDough()
	s := c.ingredientFactory.CreateSauce()
	ch := c.ingredientFactory.CreateCheese()
	fmt.Printf("Laying on the ingredients: %s, %s, %s\n", d.Name(), s.Name(), ch.Name())
}

func (c cheese) Bake() {
}

func (c cheese) Cut() {
}

func (c cheese) Box() {
}


