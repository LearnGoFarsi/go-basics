package main

import (
	"fmt"
)

type Mutex struct {
}

func (c Mutex) Lock() {
	fmt.Println("Mutex is locked")
}

func (c Mutex) Unlock() {
	fmt.Println("Mutex is unlocked")
}

type Cache struct {
	Mutex // anonymous field
	Store map[string]interface{}
}

type Shadower struct {
	Mutex // anonymous field
	Store map[string]interface{}
}

// This method will shadow the Mutex.Lock() method
func (s *Shadower) Lock() {
	fmt.Println("Shadower shadowed the Lock method of the Mutex")
}

func main() {

	c := &Cache{
		Store: make(map[string]interface{}),
	}
	c.Lock() // Check the stdout
	c.Store["Lang"] = "Go"
	c.Unlock()

	///////////////

	s := &Shadower{
		Store: make(map[string]interface{}),
	}
	s.Lock() // Check the stdout
	s.Store["Lang"] = "Go"
	s.Unlock()
}
