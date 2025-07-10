package main

import "fmt"

// user represents a system user type.
type user struct {
	name  string
	email string
	ext   int
	admin bool
}

type admin struct {
	person user // It´s not type embedding
	level  string
}

func main() {

	// Declare a user using struct literal.
	// Field order does not matter.
	marcp := user{
		admin: true,
		email: "marcpires@gmail.com",
		name:  "Marc Pires",
		ext:   123,
	}

	// Declaring a user using just the values.
	// Order matters
	bruno := user{"bruno@lhc.net.br", "Bruno", 223, false}

	douglas := admin{
		person: user{
			name:  "Douglas",
			email: "douglas@lhc.net",
			ext:   334,
			admin: true,
		},
		level: "super",
	}

	fmt.Println(marcp)
	fmt.Println(bruno)
	fmt.Println(douglas)
}
