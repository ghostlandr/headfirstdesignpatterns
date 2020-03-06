package drinks

import (
	"fmt"
)

// NewTea creates a new Drink, modelled after a Tea
func NewTea() Drink {
	return &tea{}
}

type tea struct {
	addCondiments bool
}

func (c *tea) Brew() {
	fmt.Println("Inserting tea bag")
}

func (c *tea) AddCondiments() {
	if c.addCondiments {
		fmt.Println("Adding cream and sugar")
	} else {
		fmt.Println("Skipping condiments")
	}
}
