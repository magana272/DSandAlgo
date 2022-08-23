package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type rectangle struct {
	height float64
	width  float64
}
type Shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r rectangle) area() float64 {
	return r.height * r.width
}

func main() {
	fmt.Println("Hello")
	myShapes := []Shape{rectangle{6, 4}, circle{2}}
	for _, shape := range myShapes {
		fmt.Println(shape.area())
	}

}
