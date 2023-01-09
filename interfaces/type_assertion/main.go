package main

import "fmt"

type MyStruct struct{}

func main() {
	Generic("string")
	Generic(123)
	Generic(MyStruct{})
	Generic(1.0)
}

func Generic(a interface{}) {

	switch a.(type) {
	case int:
		fmt.Println("I am an int type")
	case string:
		fmt.Println("I am a string type")
	case MyStruct:
		fmt.Println("I am a MyStruct type")
	default:
		fmt.Println("Unknown")
	}
}
