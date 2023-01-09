package main

import (
	"fmt"
)

type MyError struct {
	msg   string
	stack string
}

func (e MyError) Error() string {
	return e.msg
}

func main() {
	getError(MyError{
		msg:   "I am an error",
		stack: "stack",
	})
}

func getError(e error) {
	fmt.Println(e.Error())
}
