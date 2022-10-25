package main

import "fmt"

func main() {
	s := sum("Go", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println("Sum:", s)
}

func sum(s string, nums ...int) (r int) {
	fmt.Println("My first arg is:", s)
	for _, num := range nums {
		r += num
	}
	return
}
