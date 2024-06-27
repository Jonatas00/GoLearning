package main

import (
	"fmt"
	"math"
)

type retangle struct {
	width  float64
	height float64
}

func (r retangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type model interface {
	area() float64
}

func printArea(m model) {
	fmt.Printf("The area of is %f \n", m.area())
}

func main() {
	r1 := retangle{width: 10, height: 10}
	fmt.Println(r1.area())

	c1 := circle{radius: 5}
	printArea(c1)

	printArea(r1)

}
