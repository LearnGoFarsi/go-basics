package main

import (
	"fmt"
)

func squares() func() {
	x := 0

	return func() {
		// closure
		x++
		fmt.Println(x * x)
	}
}

func loop() {
	for i := 0; i < 10; i++ {
		func(i int) {
			i *= i
			fmt.Println(i)
		}(i)
	}
}

func recursion() {
	// fibo sequence: 0 1 1 2

	var fibo func(i, j int)

	fibo = func(i, j int) {
		if i > 30 {
			fmt.Println()
			return
		}

		fmt.Printf("%d ", i)
		fibo(j, j+i)
	}

	fibo(0, 1)
}

func main() {
	// sqaure
	f := squares()
	f()
	f()
	squares()()

	// loop
	loop()

	// fibo
	recursion()
}
