package synchronized

import (
	"fmt"
	"sync"
)

type Thinger interface {
	Thing()
}

var mu sync.Mutex
var t *thing

func GetInstance() Thinger {
	mu.Lock()
	defer mu.Unlock()
	if t == nil {
		t = &thing{expensiveDBConnection: "so expensive"}
	}
	return t
}

type thing struct {
	expensiveDBConnection string
}

func (t *thing) Thing() {
	fmt.Println("Whatever exactly this is supposed to do ...")
}
