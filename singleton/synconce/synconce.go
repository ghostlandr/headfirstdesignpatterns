package synconce

import (
	"fmt"
	"sync"

	"headfirstdesigntraining/singleton/synchronized"
)

var t *thing
var initOnce sync.Once

func GetInstance() synchronized.Thinger {
	initOnce.Do(func() {
		t = &thing{expensiveDBConnection: "So expensive"}
	})
	return t
}

type thing struct {
	expensiveDBConnection string
}

func (t *thing) Thing() {
	fmt.Println("Whatever exactly this is supposed to do ...")
}
