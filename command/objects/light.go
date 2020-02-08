package objects

import "fmt"

type Light struct {}

func (l Light) On() {
	fmt.Println("Turning light on")
}

func (l Light) Off() {
	fmt.Println("Turning light off")
}

