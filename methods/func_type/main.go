package main

import "fmt"

type Authorized func(age int) bool

// A func type can also have some methods.
// It is a very rare case with some special use cases

// The Alcohol method has a fixed impl. so behaves always the same.
func (f Authorized) Alcohol(age int) {
	if f(age) {
		fmt.Println("You are allowed to drink Alcohol!")
		// some code
	} else {
		fmt.Println("Alcohol is not allowed for you!")
		// some code
	}
}

func main() {

	// The Authorized func has a dynamic impl.
	loose := func(age int) bool {
		fmt.Println("I am a loose func.")
		return age > 17
	}

	// The Authorized func has a dynamic impl.
	tight := func(age int) bool {
		fmt.Println("I am a tight func.")
		return age > 23
	}

	Authorized(loose).Alcohol(18)
	Authorized(tight).Alcohol(18)
}
