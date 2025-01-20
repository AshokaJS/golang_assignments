package main

import (
	"fmt"
)

type rectangle struct {
	length  int
	breadth int
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func (r rectangle) perimeter() int {
	return 2 * (r.breadth + r.length)
}

func main() {
	var length, breadth int
	fmt.Scanln(&length)
	fmt.Scanln(&breadth)

	var r rectangle
	r = rectangle{length, breadth}
	fmt.Println("area of rectangle: ", r.area())

	fmt.Println("perimeter of rectangle : ", r.perimeter())

}
