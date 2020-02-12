package drinks

import (
	"fmt"
)

func NewCoffee() Drink {
	return &coffee{}
}

type coffee struct {
	addCondiments bool
}

func (c *coffee) Brew() {
	fmt.Println("Brewing coffee")
}

func (c *coffee) AddCondiments() {
	if c.addCondiments {
		fmt.Println("Adding cream and sugar")
	} else {
		fmt.Println("Skipping condiments")
	}
}
