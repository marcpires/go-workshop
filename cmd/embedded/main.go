package main

import (
	"fmt"

	"github.com/marcpires/rss/pkg/inmem"
)

// Foo represents something.
type Foo struct {
	Bar
}

// Bar represents something.
// Bar field is promoted to Foo
type Bar struct {
	Baz int
}

// embedMisuse demonstrates a misuse of type embedding
// as Lock and UnLock methods gets promoted and is visible to
// external clients.
func embedMisuse() {
	m := inmem.New()
	m.Lock() // Should not be able do this
}

func main() {
	foo := Foo{}
	foo.Baz = 23

	fmt.Printf("foo.Baz value is %d is promoted", foo.Baz)
	fmt.Printf("foo.Bar.Baz value is %d", foo.Bar.Baz)

	// embedMisuse()
}
