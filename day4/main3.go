package main

import (
	"fmt"
)

type square struct {
	side int
}
type rectangle struct {
	length  int
	breadth int
}

func (s square) area() int {
	return s.side * s.side
}

func (s square) perimeter() int {
	return 4 * s.side
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func (r rectangle) perimeter() int {
	return 2 * (r.breadth + r.length)
}

type quadrilateral interface {
	area() int
	perimeter() int
}

func print(q quadrilateral) {
	fmt.Println("Area : ", q.area())
	fmt.Println("perimeter : ", q.perimeter())
}
func main() {

	length := 4
	breadth := 5

	side := 9

	var r rectangle = rectangle{length, breadth}
	var s square = square{side}

	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		print(r)
	case 2:
		print(s)
	}

}
