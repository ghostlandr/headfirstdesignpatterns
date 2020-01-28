package doublechecked

import (
	"fmt"
	"sync"

	"headfirstdesigntraining/singleton/synchronized"
)

var mu sync.Mutex
var t *thing

func GetInstance() synchronized.Thinger {
	if t == nil {
		mu.Lock()
		defer mu.Unlock()
		t = &thing{expensiveDBConnection: "Doing this on demand, protected by a single mutex call"}
	}
	return t
}

type thing struct {
	expensiveDBConnection string
}

func (t *thing) Thing() {
	fmt.Println("Whatever exactly this is supposed to do ...")
}
