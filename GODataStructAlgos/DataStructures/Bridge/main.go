//Decouples an abstraction from its implementaion so that the two can vary indpendently
//Progressively adding functionalty while separting out major difference using abstract classes

//  When should you use this pattern ?
//  - When you want to be able change both the aabstractions and concrete classes indepently

package main

import (
	"fmt"
)

// IDraw interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

// DrawShape Struct
type DrawShape struct {
}

// drawShape method
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {

	fmt.Println("Drawing Shape")
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

// DrawContour struct
type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

// DrawContour method
func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

// DrawContour method resizeByFactor given factor
func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

// main
func main() {
	fmt.Println("Bridge strcutal pattern ")
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
