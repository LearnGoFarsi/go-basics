package main

import "fmt"

type IntList struct {
	Val  int
	Next *IntList
}

func (i *IntList) Sum() int {
	// Nil is a valid reciever of any *Type, so we need to handle it properly.
	if i == nil {
		return 0
	}

	return i.Val + i.Next.Sum()
}

func main() {

	list := &IntList{
		Val: 12,
		Next: &IntList{
			Val: 4,
			Next: &IntList{
				Val: 6,
			},
		},
	}

	fmt.Println(list.Sum())
}
