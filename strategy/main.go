// Implementing the duck examples from chapter 1 of Headfirst Design Patterns:
// The Strategy pattern
package main

import (
	"fmt"
)

type Flyer interface {
	Fly()
}

type Quacker interface {
	Quack()
}

type Duck struct {
	fly Flyer
	quack Quacker
}

func (d Duck) Quack() {
	d.getQuacker().Quack()
}

func (d Duck) Fly() {
	d.getFlyer().Fly()
}

func (d Duck) getQuacker() Quacker {
	return d.quack
}

func (d Duck) getFlyer() Flyer {
	return d.fly
}

func NewDuck(q Quacker, f Flyer) Duck {
	return Duck{fly: f, quack: q}
}

func (d *Duck) SetQuacker(q Quacker) {
	d.quack = q
}

func (d *Duck) SetFlyer(f Flyer) {
	d.fly = f
}

type ItQuacks struct{}
func (q ItQuacks) Quack() {
	fmt.Println("We're quacking!")
}

func NewItQuacks() Quacker {
	return ItQuacks{}
}

type SilentQuack struct{}
func (q SilentQuack) Quack() {
	fmt.Println("We don't quack")
}

func NewSilentQuack() Quacker {
	return SilentQuack{}
}

type FlyWithWings struct{}
func (f FlyWithWings) Fly() {
	fmt.Println("Flying with wings!")
}

func NewFlyWithWings() Flyer {
	return FlyWithWings{}
}

type FlyNoWings struct{}
func (f FlyNoWings) Fly() {
	fmt.Println("We can't fly, no wings :(")
}

func NewFlyNoWings() Flyer {
	return FlyNoWings{}
}

type FlyRocketPowered struct{}
func (f FlyRocketPowered) Fly() {
	fmt.Println("Flying with rocket power!")
}

func NewFlyRocketPowered() Flyer {
	return FlyRocketPowered{}
}

func main() {
	mallard := NewDuck(NewItQuacks(), NewFlyWithWings())
	mallard.Quack()
	mallard.Fly()

	rubber := NewDuck(NewItQuacks(), NewFlyNoWings())
	rubber.Quack()
	rubber.Fly()

	woodenDuck := NewDuck(NewSilentQuack(), NewFlyNoWings())
	woodenDuck.Quack()
	woodenDuck.Fly()

	modelDuck := NewDuck(NewItQuacks(), NewFlyNoWings())
	modelDuck.Quack()
	modelDuck.Fly() // No flying :(
	modelDuck.SetFlyer(NewFlyRocketPowered())
	modelDuck.Fly() // We flyin'!
}

