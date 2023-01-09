package main

import "fmt"

// /// EXTERNAL CODE
type ExternalStruct struct{}

func (e *ExternalStruct) MyMethod() error {
	// Implementation
	fmt.Println("I am the original method")
	return nil
}

/////

type MyMock interface {
	MyMethod() error
}

type MockStruct struct{}

func (e *MockStruct) MyMethod() error {
	// Implementation
	fmt.Println("I am a mock method")
	return nil
}

func Call(e MyMock) {
	e.MyMethod()
}

func main() {
	Call(&ExternalStruct{})

	Call(&MockStruct{})
}
