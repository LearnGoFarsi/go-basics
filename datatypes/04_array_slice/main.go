package main

import "fmt"

func main() {

	// array is a fixed-length sequence of zero or more elements of a particular type

	a := [10]string{} // a is initialied with the zero value of its type
	fmt.Println("array a:", a, len(a), cap(a))

	for i := 0; i < 5; i++ {
		a[i] = "a"
	}
	fmt.Println("array a:", a, len(a), cap(a))
	fmt.Println("Array type is comparable, right?", [10]string{} == [10]string{})

	// index 10 out of bound
	// a[10] = "a"

	// list of index and value pair
	// source 'The Go programing language'
	type number int
	const (
		zero number = iota
		one
		two
		three
	)
	a2 := [...]string{zero: "zero", one: "one", two: "two", three: "three"}
	fmt.Println("two is:", two, a2[two])

	a3 := [...]string{5: "a"} // other indexes are intialized with zero value of type string
	fmt.Println("a3 is:", a3, len(a3), cap(a3))

	//////////////////////////////////

	// slice is a variable-length sequence of elements of a particular type
	// allocates an underlying array of size 0 and
	// returns a slice of length 0 and capacity 0 that is backed by this underlying array.
	s := []string{}
	fmt.Println("slice s:", s, len(s), cap(s))

	// index out of range [0] with length 0
	//for i := 0; i < 5; i++ {
	//	s[i] = "s"
	//}
	s = append(s, "s")
	fmt.Println("slice s:", s, len(s), cap(s))

	// Slice type is not comparable
	//fmt.Println([]string{"a"} == []string{"v"})

	//////////////////////////////////

	// allocates an underlying array of size 10 and
	// returns a slice of length 0 and capacity 10 that is backed by this underlying array.
	s2 := make([]string, 0, 10)
	fmt.Println("slice s2:", s2, len(s2), cap(s2))

	// index 0 out of bound
	//s2[0] = "Go"

	// Will NOT run out of the range
	for i := 0; i < 15; i++ {
		s2 = append(s2, "a")
	}
	fmt.Println("slice s2:", s2, len(s2), cap(s2)) // Notice: capacity grew beyond 15

	//////////////////////////////////

	s3 := []string{"a", "b", "c", "d"}
	s4 := s3[:2] // Python like
	s4[1] = "B"
	fmt.Println(s3) // s3 and s4 share the same underlying array
}
