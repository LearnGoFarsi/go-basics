package main

import (
	"fmt"
	"time"
)

type Rocket struct {
}

func (r *Rocket) Launch() {
	fmt.Println("Rocket Launched")
}

func Delayed(d time.Duration, f func()) {
	fmt.Printf("Rocket will be launched in %d seconds.\n", d/1000000000)
	<-time.After(d)
	f()
}

func main() {

	rocket := Rocket{}
	// launch is a method value
	launch := rocket.Launch
	Delayed(time.Second*3, launch)
	// or
	Delayed(time.Second*3, rocket.Launch)
}
