package main

import "fmt"

type duck interface {
	Quack()
	Fly()
}

type turkey interface {
	Gobble()
	FlyShortDistance()
}

type turkeyAdapter struct {
	t turkey
}

func (t *turkeyAdapter) Quack() {
	t.t.Gobble()
}

func (t *turkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.t.FlyShortDistance()
	}
}

type turkeyImpl struct {}
func (t *turkeyImpl) Gobble() { fmt.Println("Gobble, gobble") }
func (t *turkeyImpl) FlyShortDistance() { fmt.Println("I'm flying a short distance") }

func testDuck(d duck) {
	d.Quack()
	d.Fly()
}

func main() {
	turk := &turkeyImpl{}
	durk := &turkeyAdapter{t: turk}
	testDuck(durk)
}
