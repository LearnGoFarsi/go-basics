package main

import "fmt"

func main() {
	// a struct is an aggregate data type that groups zero or more named values of arbitrary types
	// each value is called a field

	type Person struct {
		Name, Address string
	}

	type Order struct {
		ID        int
		ItemName  string
		ItemValue float32
	}

	type Employee struct {
		Person
		Salary int
		Orders []*Order
	}

	// pMilad is a pointer
	pMilad := &Employee{
		Person: Person{
			Name:    "Milad",
			Address: "Planet Earth",
		},
		Salary: 100,
	}

	pMilad.Orders = append(pMilad.Orders, &Order{
		ID:        1,
		ItemName:  "Icecream",
		ItemValue: 2,
	})

	// Copy value of the Milad's pointer
	// UnnamedEmp is not a pointer
	UnnamedEmp := *pMilad

	// A struct is comparable only if all its fields are comparable
	//fmt.Println("Is UnnameEmp equal to Milad?", UnnamedEmp == *pMilad)

	// These are two different variables with the same content
	UnnamedEmp.Name = "Go"
	UnnamedEmp.Orders[0].ID = 987
	fmt.Println(UnnamedEmp.Name, pMilad.Name)
	fmt.Println(UnnamedEmp.Orders[0], pMilad.Orders[0]) // Ohhhhhhh

	tree := &BinaryTree{
		Value: 0,
	}

	tree.Add(3)
	tree.Add(1)
	tree.Add(7)
	tree.Add(4)

	fmt.Println("Can you find 3 in the tree?", tree.Has(3))
	fmt.Println("Can you find 9 in the tree?", tree.Has(9))
}

type BinaryTree struct {
	right *BinaryTree
	left  *BinaryTree
	Value int
}

func (t *BinaryTree) Add(val int) {
	t.add(val)
}

func (t *BinaryTree) add(val int) {
	if val > t.Value {
		if t.right == nil {
			t.right = &BinaryTree{Value: val}
		} else {
			t.right.add(val)
		}
	} else if val <= t.Value {
		if t.left == nil {
			t.left = &BinaryTree{Value: val}
		} else {
			t.left.add(val)
		}
	}
}

func (t *BinaryTree) Has(val int) (b bool) {
	if t.Value == val {
		b = true
	} else if t.Value < val && t.right != nil {
		b = t.right.Has(val)
	} else if t.Value > val && t.left != nil {
		b = t.left.Has(val)
	} else {
		b = false
	}

	return
}
