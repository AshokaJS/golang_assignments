// code for calculating area of circle in golang

package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Scanln(&radius)

	var area float64

	area = math.Pow(radius, 2)

	area = math.Pi * area
	roundedArea := fmt.Sprintf("%.2f", area)
	fmt.Println("Area of circle is :", roundedArea)

}
