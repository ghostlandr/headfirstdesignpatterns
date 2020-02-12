package brewer

import (
	"fmt"

	"headfirstdesigntraining/template/drinks"
)

type Beverage interface {
	PrepareRecipe()
	BoilWater()
	PutInCup()
	Brew()
	AddCondiments()
}

func NewBeverage(d drinks.Drink) Beverage {
	return &caffeineBeverage{Drink: d}
}

type caffeineBeverage struct {
	// Leaving it to Drink to define Brew and AddCondiments
	drinks.Drink
}

func (c *caffeineBeverage) PrepareRecipe() {
	c.BoilWater()
	c.PutInCup()
	c.Brew()
	c.AddCondiments()
}

func (c *caffeineBeverage) BoilWater() {
	fmt.Println("Getting water up to 200 degrees F")
}

func (c *caffeineBeverage) PutInCup() {
	fmt.Println("Pouring into a cup")
}
