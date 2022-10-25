package main

import (
	"fmt"
	"time"
)

type Revolution struct {
	Name  string
	Date  int
	Heros []string
}

func main() {
	intro()

	r := Revolution{
		Name: "Zan, Zendegi, Azadi",
		Date: time.Now().Year(),
	}

	passByValue(r)
	fmt.Println(r)

	passByRef(&r)
	fmt.Println(r)
}

func intro() {
	var i int8 = 54
	fmt.Println("memory address of i is:", &i)

	var p *int8
	p = &i
	fmt.Printf("p points to %p memory address: \n", p)
	fmt.Printf("memory address %p stores value %d \n", p, *p)

	*p = 100
	fmt.Println("Value of i is: ", i)
}

func passByValue(r Revolution) {
	r.Heros = append(r.Heros, "Mahsa", "Khodanoor", "Kian", "80.000.000+")
}

func passByRef(r *Revolution) {
	r.Heros = append(r.Heros, "Mahsa", "Khodanoor", "Kian", "80.000.000+")
}
