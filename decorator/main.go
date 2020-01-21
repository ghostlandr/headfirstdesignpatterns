// Implementing the pattern from chapter 3 (Decorator Pattern)
package main

import (
	"fmt"

	"headfirstdesigntraining/decorator/beverage"
)

func main() {
	e := beverage.Espresso()
	fmt.Printf("%s: $%.2f\n", e.Description(), e.Cost())

	// Add some mocha
	e = beverage.Mocha(e)
	fmt.Printf("%s: $%.2f\n", e.Description(), e.Cost())

	// Add soy
	e = beverage.Whip(e)
	fmt.Printf("%s: $%.2f\n", e.Description(), e.Cost())

	// It's a big one (size-related pricing only implemented for soy)
	e.SetSize(beverage.Venti)
	fmt.Printf("%s: $%.2f\n", e.Description(), e.Cost())

	// Add some whip
	e = beverage.Whip(e)
	fmt.Printf("%s: $%.2f\n", e.Description(), e.Cost())
}
