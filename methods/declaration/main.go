package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Distance is a method of *Point. We can access this method only through a *Point type
func (p *Point) Distance(q Point) float64 {
	if p == nil {
		return 0
	}

	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

// MoveBy is a method of *Point. We can access this method only through a *Point type
func (p *Point) MoveBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []*Point

// Distance is a method of Path. We can access this method only through a Path type
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			// Each element of Path is a *Point, so we can access the (*Point).Distance()
			// method of them.
			sum += p[i-1].Distance(*p[i])
		}
	}

	return sum
}

func main() {

	p := &Point{
		X: 5, Y: 7,
	}
	p.MoveBy(2)
	fmt.Println(p)

	q := Point{X: 6, Y: 2}
	// Distance of p from q
	p.Distance(q)

	path := Path{
		{3, 5}, {6, 1}, {9, 2},
	}
	path.Distance()
}
