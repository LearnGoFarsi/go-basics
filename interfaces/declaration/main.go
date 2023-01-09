package main

import "fmt"

type MyInterface interface {
	Hello(string)
	Bye(string) error
}

type MyStruct struct{}

// To make sure that MyStruct impl. MyInterface
var _ MyInterface = (*MyStruct)(nil)

func (s *MyStruct) Hello(name string) {
	fmt.Printf("Hello %s! I am MyStruct \n", name)
}

func (s *MyStruct) Bye(name string) error {
	fmt.Printf("Bye %s \n", name)
	return nil
}

type MyInt int

func (s *MyInt) Hello(name string) {
	fmt.Printf("Hello %s! I am MyInt \n", name)
}

func (s MyInt) Bye(name string) error {
	fmt.Printf("Bye %s \n", name)
	return nil
}

func main() {

	myStruct := &MyStruct{}
	myStruct.Hello("Iran")
	myStruct.Bye("Islamic republic")

	myInt := MyInt(10)
	myInt.Hello("Mahsa")
	myInt.Bye("Akhoonds")

	var myInterface MyInterface
	// Because MyStruct satisfies MyInterface
	myInterface = &MyStruct{}
	myInterface.Hello("Democracy")
	myInterface.Bye("Dictator")
}
