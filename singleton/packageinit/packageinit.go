package packageinit

import (
	"fmt"

	"headfirstdesigntraining/singleton/synchronized"
)

type thing struct {
	expensiveDBConnection string
}

func (t *thing) Thing() {
	fmt.Println("Whatever exactly this is supposed to do ...")
}

var t *thing

func init() {
	t = &thing{expensiveDBConnection: "So expensive it's good to do it up front"}
}

func GetInstance() synchronized.Thinger {
	return t
}
