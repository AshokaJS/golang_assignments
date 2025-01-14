// code for calculating simple interest in golang.
package main

import "fmt"

func main() {
	var p int
	var r int
	var t int

	fmt.Scanln(&p)
	fmt.Scanln(&r)
	fmt.Scanln(&t)
	
	var simple_interest int
	simple_interest = (p * r * t) / 100

	fmt.Println(simple_interest)
}
