package main

import (
	"fmt"

	"github.com/teq0/musical-umbrella/pkg/magic"
)

func main() {
	t := &magic.Thing{Name: "thing"}
	fmt.Println(t.Name)
	t.DoThing()
}
