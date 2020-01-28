package main

import (
	"headfirstdesigntraining/singleton/doublechecked"
	"headfirstdesigntraining/singleton/packageinit"
	"headfirstdesigntraining/singleton/synchronized"
	"headfirstdesigntraining/singleton/synconce"
)

func main() {
	obj1 := synchronized.GetInstance()
	obj2 := packageinit.GetInstance()
	obj3 := doublechecked.GetInstance()
	obj4 := synconce.GetInstance()

	obj1.Thing()
	obj2.Thing()
	obj3.Thing()
	obj4.Thing()
}
