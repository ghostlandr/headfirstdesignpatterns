package main

import (
	"headfirstdesigntraining/template/brewer"
	"headfirstdesigntraining/template/drinks"
)

func main() {
	c := brewer.NewBeverage(drinks.NewCoffee())
	c.PrepareRecipe()

	t := brewer.NewBeverage(drinks.NewTea())
	t.PrepareRecipe()

	g := brewer.NewBeverage(drinks.NewTea())
	g.PrepareRecipe()
}
