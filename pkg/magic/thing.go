package magic

import (
	"fmt"
)

// Thing is a thing.
type Thing struct {
	Name string
}

// DoThing does a thing.

func (t *Thing) DoThing() {
	fmt.Println("Doing the thing")
}
